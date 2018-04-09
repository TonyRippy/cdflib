package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"time"

	cdflib "github.com/TonyRippy/cdflib/lang"
	zmq "github.com/pebbe/zmq4"
)

// ExecCounter is incremented each time we run user code in the notebook.
var ExecCounter int

// ConnectionInfo stores the contents of the kernel connection
// file created by Jupyter.
type ConnectionInfo struct {
	SignatureScheme string `json:"signature_scheme"`
	Transport       string `json:"transport"`
	StdinPort       int    `json:"stdin_port"`
	ControlPort     int    `json:"control_port"`
	IOPubPort       int    `json:"iopub_port"`
	HBPort          int    `json:"hb_port"`
	ShellPort       int    `json:"shell_port"`
	Key             string `json:"key"`
	IP              string `json:"ip"`
}

// SocketGroup holds the sockets needed to communicate with the kernel,
// and the key for message signing.
type SocketGroup struct {
	ShellSocket   *zmq.Socket
	ControlSocket *zmq.Socket
	StdinSocket   *zmq.Socket
	IOPubSocket   *zmq.Socket
	HBSocket      *zmq.Socket
	Key           []byte
}

// KernelLanguageInfo holds information about the language that this kernel executes code in.
type kernelLanguageInfo struct {
	Name              string `json:"name"`
	Version           string `json:"version"`
	MIMEType          string `json:"mimetype"`
	FileExtension     string `json:"file_extension"`
	PygmentsLexer     string `json:"pygments_lexer"`
	CodeMirrorMode    string `json:"codemirror_mode"`
	NBConvertExporter string `json:"nbconvert_exporter"`
}

// HelpLink stores data to be displayed in the help menu of the notebook.
type helpLink struct {
	Text string `json:"text"`
	URL  string `json:"url"`
}

// KernelInfo holds information about the igo kernel, for kernel_info_reply messages.
type kernelInfo struct {
	ProtocolVersion       string             `json:"protocol_version"`
	Implementation        string             `json:"implementation"`
	ImplementationVersion string             `json:"implementation_version"`
	LanguageInfo          kernelLanguageInfo `json:"language_info"`
	Banner                string             `json:"banner"`
	HelpLinks             []helpLink         `json:"help_links"`
}

// shutdownReply encodes a boolean indication of shutdown/restart.
type shutdownReply struct {
	Restart bool `json:"restart"`
}

const (
	kernelStarting = "starting"
	kernelBusy     = "busy"
	kernelIdle     = "idle"
)

// runKernel is the main entry point to start the kernel.
func runKernel(connectionFile string) {

	// Set up a new language environment.
	interp := cdflib.MakeInterpreter()

	// Parse the connection info.
	var connInfo ConnectionInfo

	connData, err := ioutil.ReadFile(connectionFile)
	if err != nil {
		log.Fatal(err)
	}

	if err = json.Unmarshal(connData, &connInfo); err != nil {
		log.Fatal(err)
	}

	// Set up the ZMQ sockets through which the kernel will communicate.
	sockets, err := prepareSockets(connInfo)
	if err != nil {
		log.Fatal(err)
	}

	// TODO connect all channel handlers to a WaitGroup to ensure shutdown before returning from runKernel.

	// Start up the heartbeat handler.
	startHeartbeat(sockets.HBSocket, &sync.WaitGroup{})

	// TODO gracefully shutdown the heartbeat handler on kernel shutdown by closing the chan returned by startHeartbeat.

	poller := zmq.NewPoller()
	poller.Add(sockets.ShellSocket, zmq.POLLIN)
	poller.Add(sockets.StdinSocket, zmq.POLLIN)
	poller.Add(sockets.ControlSocket, zmq.POLLIN)

	// msgParts will store a received multipart message.
	var msgParts [][]byte

	// Start a message receiving loop.
	for {
		polled, err := poller.Poll(-1)
		if err != nil {
			log.Fatal(err)
		}

		for _, item := range polled {

			// Handle various types of messages.
			switch socket := item.Socket; socket {

			// Handle shell messages.
			case sockets.ShellSocket:
				msgParts, err = sockets.ShellSocket.RecvMessageBytes(0)
				if err != nil {
					log.Println(err)
				}

				msg, ids, err := WireMsgToComposedMsg(msgParts, sockets.Key)
				if err != nil {
					log.Println(err)
					return
				}

				handleShellMsg(interp, msgReceipt{msg, ids, sockets})

				// TODO Handle stdin socket.
			case sockets.StdinSocket:
				sockets.StdinSocket.RecvMessageBytes(0)

				// Handle control messages.
			case sockets.ControlSocket:
				msgParts, err = sockets.ControlSocket.RecvMessageBytes(0)
				if err != nil {
					log.Println(err)
					return
				}

				msg, ids, err := WireMsgToComposedMsg(msgParts, sockets.Key)
				if err != nil {
					log.Println(err)
					return
				}

				handleShellMsg(interp, msgReceipt{msg, ids, sockets})
			}
		}
	}
}

// prepareSockets sets up the ZMQ sockets through which the kernel
// will communicate.
func prepareSockets(connInfo ConnectionInfo) (SocketGroup, error) {

	// Initialize the context.
	context, err := zmq.NewContext()
	if err != nil {
		return SocketGroup{}, err
	}

	// Initialize the socket group.
	var sg SocketGroup

	// Create the shell socket, a request-reply socket that may receive messages from multiple frontend for
	// code execution, introspection, auto-completion, etc.
	sg.ShellSocket, err = context.NewSocket(zmq.ROUTER)
	if err != nil {
		return sg, err
	}

	// Create the control socket. This socket is a duplicate of the shell socket where messages on this channel
	// should jump ahead of queued messages on the shell socket.
	sg.ControlSocket, err = context.NewSocket(zmq.ROUTER)
	if err != nil {
		return sg, err
	}

	// Create the stdin socket, a request-reply socket used to request user input from a front-end. This is analogous
	// to a standard input stream.
	sg.StdinSocket, err = context.NewSocket(zmq.ROUTER)
	if err != nil {
		return sg, err
	}

	// Create the iopub socket, a publisher for broadcasting data like stdout/stderr output, displaying execution
	// results or errors, kernel status, etc. to connected subscribers.
	sg.IOPubSocket, err = context.NewSocket(zmq.PUB)
	if err != nil {
		return sg, err
	}

	// Create the heartbeat socket, a request-reply socket that only allows alternating recv-send (request-reply)
	// calls. It should echo the byte strings it receives to let the requester know the kernel is still alive.
	sg.HBSocket, err = context.NewSocket(zmq.REP)
	if err != nil {
		return sg, err
	}

	// Bind the sockets.
	address := fmt.Sprintf("%v://%v:%%v", connInfo.Transport, connInfo.IP)
	sg.ShellSocket.Bind(fmt.Sprintf(address, connInfo.ShellPort))
	sg.ControlSocket.Bind(fmt.Sprintf(address, connInfo.ControlPort))
	sg.StdinSocket.Bind(fmt.Sprintf(address, connInfo.StdinPort))
	sg.IOPubSocket.Bind(fmt.Sprintf(address, connInfo.IOPubPort))
	sg.HBSocket.Bind(fmt.Sprintf(address, connInfo.HBPort))

	// Set the message signing key.
	sg.Key = []byte(connInfo.Key)

	return sg, nil
}

// handleShellMsg responds to a message on the shell ROUTER socket.
func handleShellMsg(interp cdflib.Interpreter, receipt msgReceipt) {
	switch receipt.Msg.Header.MsgType {
	case "kernel_info_request":
		if err := sendKernelInfo(receipt); err != nil {
			log.Fatal(err)
		}
	case "execute_request":
		if err := handleExecuteRequest(interp, receipt); err != nil {
			log.Fatal(err)
		}
	case "shutdown_request":
		handleShutdownRequest(receipt)
	default:
		log.Println("Unhandled shell message: ", receipt.Msg.Header.MsgType)
	}
}

// sendKernelInfo sends a kernel_info_reply message.
func sendKernelInfo(receipt msgReceipt) error {
	return receipt.Reply("kernel_info_reply",
		kernelInfo{
			ProtocolVersion:       ProtocolVersion,
			Implementation:        Name,
			ImplementationVersion: Version,
			Banner:                fmt.Sprintf("%s kernel: v%s", Name, Version),
			LanguageInfo: kernelLanguageInfo{
				Name:          Name,
				Version:       Version,
				FileExtension: FileExtension,
			},
			HelpLinks: []helpLink{
				//{Text: "Go", URL: "https://golang.org/"},
				//{Text: "gophernotes", URL: "https://github.com/gopherdata/gophernotes"},
			},
		},
	)
}

// handleExecuteRequest runs code from an execute_request method,
// and sends the various reply messages.
func handleExecuteRequest(interp cdflib.Interpreter, receipt msgReceipt) error {

	// Extract the data from the request.
	reqcontent := receipt.Msg.Content.(map[string]interface{})
	code := reqcontent["code"].(string)
	silent := reqcontent["silent"].(bool)

	if !silent {
		ExecCounter++
	}

	// Prepare the map that will hold the reply content.
	content := make(map[string]interface{})
	content["execution_count"] = ExecCounter

	// Tell the front-end that the kernel is working and when finished notify the
	// front-end that the kernel is idle again.
	if err := receipt.PublishKernelStatus(kernelBusy); err != nil {
		log.Printf("Error publishing kernel status 'busy': %v\n", err)
	}
	defer func() {
		if err := receipt.PublishKernelStatus(kernelIdle); err != nil {
			log.Printf("Error publishing kernel status 'idle': %v\n", err)
		}
	}()

	// Tell the front-end what the kernel is about to execute.
	if err := receipt.PublishExecutionInput(ExecCounter, code); err != nil {
		log.Printf("Error publishing execution input: %v\n", err)
	}

	// Create pipes for the interpreter's stdin/stderr.
	rOut, wOut, err := os.Pipe()
	if err != nil {
		return err
	}
	rErr, wErr, err := os.Pipe()
	if err != nil {
		return err
	}

	var writersWG sync.WaitGroup
	writersWG.Add(2)

	// Forward all data written to stdout/stderr to the front-end.
	go func() {
		defer writersWG.Done()
		jupyterStdOut := JupyterStreamWriter{StreamStdout, &receipt}
		io.Copy(&jupyterStdOut, rOut)
	}()

	go func() {
		defer writersWG.Done()
		jupyterStdErr := JupyterStreamWriter{StreamStderr, &receipt}
		io.Copy(&jupyterStdErr, rErr)
	}()

	oldOut := interp.SetOut(wOut)
	oldErr := interp.SetErr(wErr)
	data, executionErr := interp.Eval(code)
	interp.SetOut(oldOut)
	interp.SetErr(oldErr)

	// Close and restore the streams.
	wOut.Close()
	wErr.Close()

	// Wait for the writers to finish forwarding the data.
	writersWG.Wait()

	if executionErr == nil {
		content["status"] = "ok"
		content["user_expressions"] = make(map[string]string)

		if !silent && data != nil {
			// Publish the result of the execution.
			if err := receipt.PublishExecutionResult(ExecCounter, bundledMIMEData(data.MimeData())); err != nil {
				log.Printf("Error publishing execution result: %v\n", err)
			}
		}
	} else {
		content["status"] = "error"
		content["ename"] = "ERROR"
		content["evalue"] = executionErr.Error()
		content["traceback"] = nil

		if err := receipt.PublishExecutionError(executionErr.Error(), []string{executionErr.Error()}); err != nil {
			log.Printf("Error publishing execution error: %v\n", err)
		}
	}

	// Send the output back to the notebook.
	return receipt.Reply("execute_reply", content)
}

// handleShutdownRequest sends a "shutdown" message.
func handleShutdownRequest(receipt msgReceipt) {
	content := receipt.Msg.Content.(map[string]interface{})
	restart := content["restart"].(bool)

	reply := shutdownReply{
		Restart: restart,
	}

	if err := receipt.Reply("shutdown_reply", reply); err != nil {
		log.Fatal(err)
	}

	log.Println("Shutting down in response to shutdown_request")
	os.Exit(0)
}

// startHeartbeat starts a go-routine for handling heartbeat ping messages sent over the given `hbSocket`. The `wg`'s
// `Done` method is invoked after the thread is completely shutdown. To request a shutdown the returned `shutdown` channel
// can be closed.
func startHeartbeat(hbSocket *zmq.Socket, wg *sync.WaitGroup) (shutdown chan struct{}) {
	quit := make(chan struct{})

	// Start the handler that will echo any received messages back to the sender.
	wg.Add(1)
	go func() {
		defer wg.Done()

		// Create a `Poller` to check for incoming messages.
		poller := zmq.NewPoller()
		poller.Add(hbSocket, zmq.POLLIN)

		for {
			select {
			case <-quit:
				return
			default:
				// Check for received messages waiting at most 500ms for once to arrive.
				pingEvents, err := poller.Poll(500 * time.Millisecond)
				if err != nil {
					log.Fatalf("Error polling heartbeat channel: %v\n", err)
				}

				// If there is at least 1 message waiting then echo it.
				if len(pingEvents) > 0 {
					// Read a message from the heartbeat channel as a simple byte string.
					pingMsg, err := hbSocket.RecvBytes(0)
					if err != nil {
						log.Fatalf("Error reading heartbeat ping bytes: %v\n", err)
					}

					// Send the received byte string back to let the front-end know that the kernel is alive.
					if _, err = hbSocket.SendBytes(pingMsg, 0); err != nil {
						log.Printf("Error sending heartbeat pong bytes: %b\n", err)
					}
				}
			}
		}
	}()

	return quit
}
