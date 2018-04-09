package main

import (
	"flag"
	"log"
)

const (
	// Name used to identify this kernel.
	Name string = "cdflib"
	
	// Version defines the kernel version.
	Version string = "0.0.1"

	// ProtocolVersion defines the Jupyter protocol version.
	ProtocolVersion string = "5.0"
	
	FileExtension string = ".cdf"
)

func main() {

	// Parse the connection file.
	flag.Parse()
	if flag.NArg() < 1 {
		log.Fatalln("Need a command line argument specifying the connection file.")
	}

	// Run the kernel.
	runKernel(flag.Arg(0))
}
