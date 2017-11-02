package main

import (
	"github.com/TonyRippy/cdflib"
	"image/color"
	"image/png"
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile("pdf.png", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Panicf("Failed to open file! (%s)", err)
	}
	defer f.Close()
	cdf := cdflib.Gaussian(0, 1)
	img := cdflib.DrawPDF(cdf, 0.999, color.Black, color.RGBA{0, 0, 0xff, 0xff}, 800, 600)
	err = png.Encode(f, img)
	if err != nil {
		log.Panicf("Failed to write file! (%s)", err)
	}
}
