package lang

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/TonyRippy/cdflib"
	"image"
	"image/color"
	"image/png"
	"strconv"
	"strings"
)

const (
	DefaultFillAlpha = 0.8
	DefaultWidth = 800
	DefaultHeight = 600
)

var (
	// From https://en.wikipedia.org/wiki/Web_colors
	NamedColors = map[string]*color.RGBA{
		"white": &color.RGBA{0xFF, 0xFF, 0xFF, 0xFF},
		"silver": &color.RGBA{0xC0, 0xC0, 0xC0, 0xFF},
		"gray": &color.RGBA{0x80, 0x80, 0x80, 0xFF},
		"black": &color.RGBA{0x00, 0x00, 0x00, 0xFF},
		"red": &color.RGBA{0xFF, 0x00, 0x00, 0xFF},
		"maroon": &color.RGBA{0x80, 0x00, 0x00, 0xFF},
		"yellow": &color.RGBA{0xFF, 0xFF, 0x00, 0xFF},
		"olive": &color.RGBA{0x80, 0x80, 0x00, 0xFF},
		"lime": &color.RGBA{0x00, 0xFF, 0x00, 0xFF},
		"green": &color.RGBA{0x00, 0x80, 0x00, 0xFF},
		"aqua": &color.RGBA{0x00, 0xFF, 0xFF, 0xFF},
		"teal": &color.RGBA{0x00, 0x80, 0x80, 0xFF},
		"blue": &color.RGBA{0x00, 0x00, 0xFF, 0xFF},
		"navy": &color.RGBA{0x00, 0x00, 0x80, 0xFF},
		"fuchsia": &color.RGBA{0xFF, 0x00, 0xFF, 0xFF},
		"purple": &color.RGBA{0x80, 0x00, 0x80, 0xFF},
	}

	PositionColors = []string{
		"blue",
		"red",
		"green",
		"yellow",
		"purple",
		"teal",
		"silver",
		"maroon",
	}
)

type Color interface {
	Eval(pos int) (*color.RGBA, error)
}

type namedColor struct {
	Name string
}

func (nc *namedColor) Eval(pos int) (*color.RGBA, error) {
	name := strings.ToLower(nc.Name)
	c, ok := NamedColors[name]
	if !ok {
		return nil, fmt.Errorf("Unknown color \"%s\".", name)
	}
	return c, nil
}

func ColorByName(name string) Color {
	return &namedColor{name}
}

type specColor struct {
	Spec string
}

func (sc *specColor) Eval(pos int) (*color.RGBA, error) {
	var ps rune
	var rs, gs, bs string
	ok := true
	n := len(sc.Spec)
	if n == 4 {
		ps = rune(sc.Spec[0])
		rs = sc.Spec[1:2] + "0" 
		gs = sc.Spec[2:3] + "0" 
		bs = sc.Spec[3:4] + "0" 
	} else if n == 7 {
		ps = rune(sc.Spec[0])
		rs = sc.Spec[1:3] 
		gs = sc.Spec[3:5]
		bs = sc.Spec[5:7]
	} else {
		ok = false
	}
	if ok {
		ok = (ps == '#')
	}
	var r, g, b uint8
	if ok {
		v, err := strconv.ParseInt(rs, 16, 8)
		r = uint8(v)
		ok = err != nil
	}
	if ok {
		v, err := strconv.ParseInt(gs, 16, 8)
		g = uint8(v)
		ok = err != nil
	}
	if ok {
		v, err := strconv.ParseInt(bs, 16, 8)
		b = uint8(v)
		ok = err != nil
	}
	if ok {
		return &color.RGBA{r,g,b,0xFF}, nil
	}
	return nil, fmt.Errorf("Invalid spec \"%s\". Expected #RGB or #RRGGBB.", sc.Spec)
}

func ColorBySpec(spec string) Color {
	return &specColor{spec}
}

type posColor struct {
}

func (pc *posColor) Eval(pos int) (*color.RGBA, error) {
	n := len(PositionColors)
	if pos < 0 || pos >= n {
		return nil, fmt.Errorf("Default colors only available for the first %d elements.", n)
	}
	return NamedColors[PositionColors[pos]], nil
}

func ColorByPosition() Color {
	return &posColor{}
}

type PlotArg struct {
	Value Expression
	Color Color
	FillAlpha float64
}

type imgExpr struct {
	Image image.Image
}

func (i *imgExpr) Type() ExpressionType {
	return IMAGE
}

func (i *imgExpr) Eval(env Environment) (Expression, error) {
	return i, nil
}

func (i *imgExpr) MimeData() MimeData {
	var buf bytes.Buffer
	f := base64.NewEncoder(base64.StdEncoding, &buf)
	png.Encode(f, i.Image)
	f.Close()
	return MimeData{
		"image/png": buf.String(),
	}
}

func (i *imgExpr) String() string {
	return "<IMAGE>"
}

type drawFunc func([]cdflib.CDF, float64, []color.Color, []color.Color, int, int)image.Image

type plotExpr struct {
	draw drawFunc
	args []PlotArg
}

func (i *plotExpr) Type() ExpressionType {
	return IMAGE
}

func (pe *plotExpr) Eval(env Environment) (Expression, error) {
	if pe.draw == nil {
		return nil, errors.New("Could not resolve plot function.")
	}
	// Look up the configured dimentions of output images.
	w := DefaultWidth
	v, ok := env.GetVar("__WIDTH__")
	if ok {
		var err error
		if w, err = AsInt(v); err != nil {
			return nil, fmt.Errorf("__WIDTH__: %s", err)
		}
	}
	h := DefaultHeight
	v, ok = env.GetVar("__HEIGHT__")
	if ok {
		var err error
		if h, err = AsInt(v); err != nil {
			return nil, fmt.Errorf("__HEIGHT__: %s", err)
		}
	}
	// Evaluate the plot arguments.
	n := len(pe.args)
	cdfs := make([]cdflib.CDF, n) 
	line := make([]color.Color, n) 
	fill := make([]color.Color, n)
	for i, pa := range(pe.args) {
		var v Expression
		var err error
		if v, err = pa.Value.Eval(env); err != nil {
			return nil, fmt.Errorf("Unable to evaluate plot argument %d: %s", i, err)
		}
		if cdfs[i], err = AsCDF(v); err != nil {
			return nil, fmt.Errorf("Unable to evaluate plot argument %d: %s", i, err)
		}
		var c *color.RGBA
		if c, err = pa.Color.Eval(i); err != nil {
			return nil, fmt.Errorf("Unable to evaluate color of plot argument %d: %s", i, err)
		}
		line[i] = c
		a := pa.FillAlpha
		if a < 0 { a = 0 }
		if a > 1 { a = 1 }
		a *= 255
		fill[i] = color.RGBA{c.R, c.G, c.B, uint8(a * 255)} 
	}

	// Draw the graph.
	img := pe.draw(cdfs, 0.999, line, fill, w, h)
	return &imgExpr{img}, nil
}

func (i *plotExpr) MimeData() MimeData {
	return MimeData{}
}

func (i *plotExpr) String() string {
	return "<PLOT>"
}

func Plot(fn string, args []PlotArg) Expression {
	var f drawFunc
	if fn == "Plot" {
		f = cdflib.DrawCDFs
	} else if fn == "Plotd" {
		f = cdflib.DrawPDFs
	}
	return &plotExpr{f, args}
}

func PlotNoArgs(fn string) Expression {
	return &errorExpr{IMAGE, fmt.Errorf("%s: Nothing to plot.", fn)}
}
