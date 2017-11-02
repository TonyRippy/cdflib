package cdflib

import (
	"image"
	"image/color"
	"image/draw"
)

func findRange(cdfs []CDF, p float64) (float64, float64) {
	inv := cdfs[0].Inverse()
	p2 := 1 - p
	if p > p2 {
		p, p2 = p2, p
	}
	min := inv.Value(p)
	max := inv.Value(1 - p)
	for i := 1; i < len(cdfs); i += 1 {
		inv = cdfs[i].Inverse()
		newMin := inv.Value(p)
		if newMin < min {
			min = newMin
		}
		newMax := inv.Value(1 - p)
		if newMax > max {
			max = newMax
		}
	}
	return min, max
}

/*
Plots a CDF and returns it as an image.

The Y axis of the image will always be [0, 1] inclusive.

The X axis is specified by the p parameter. It specifies the probability range
that should be plotted. For example, using p=0.05 the X axis of the image will
include the values in the range [V(.05), V(.95)] inclusive.

The background is set to the transparent value.
*/
func DrawCDF(cdf CDF, p float64, lineColor color.Color, fillColor color.Color, w int, h int) image.Image {
	return DrawCDFs([]CDF{cdf}, p, []color.Color{lineColor}, []color.Color{fillColor}, w, h)
}

/*
Plots one or more CDFs on the same graph.
See DrawCDF for details.
*/
func DrawCDFs(cdfs []CDF, p float64, lineColors []color.Color, fillColors []color.Color, w int, h int) image.Image {
	n := len(cdfs)
	if n == 0 {
		return image.NewUniform(color.Transparent)
	}

	// Find the widest range of X values:
	minX, maxX := findRange(cdfs, p)

	// Sample all cdfs.
	samples := make([][]float64, n)
	for i, cdf := range cdfs {
		samples[i] = UniformSamples(cdf, minX, maxX, w)
	}

	// For each CDF: Generate image, layer on top of existing images.
	var img *image.RGBA
	bounds := image.Rect(0, 0, w, h)
	h -= 1
	m := float64(h)
	for i, sample := range samples {
		newImg := image.NewRGBA(bounds)
		var lastY int
		for x, val := range sample {
			newY := h - int(val*m)
			var y1, y2 int
			if x == 0 {
				y1, y2 = newY, newY
			} else if newY > lastY {
				y1, y2 = lastY, newY
			} else {
				y1, y2 = newY, lastY
			}
			for y := y1; y <= y2; y += 1 {
				newImg.Set(x, y, lineColors[i])
			}
			for y := y2 + 1; y <= h; y += 1 {
				newImg.Set(x, y, fillColors[i])
			}
			lastY = newY
		}
		if img == nil {
			img = newImg
		} else {
			draw.Draw(img, bounds, newImg, image.ZP, draw.Over)
		}
	}
	return img
}

/*
Plots a PDF and returns it as an image.

The X axis is specified by the p parameter. It specifies the probability range
that should be plotted. For example, using p=0.05 the X axis of the image will
include the values in the range [V(.05), V(.95)] inclusive.

The Y axis if the image will always be in the range (0, max),
where max is the largest Y value of the PDF.

The background is set to the transparent value.
*/
func DrawPDF(cdf CDF, p float64, lineColor color.Color, fillColor color.Color, w int, h int) image.Image {
	return DrawPDFs([]CDF{cdf}, p, []color.Color{lineColor}, []color.Color{fillColor}, w, h)
}

/*
Plots one or more PDFs in the same image.
See DrawPDF fr details.
*/
func DrawPDFs(cdfs []CDF, p float64, lineColors []color.Color, fillColors []color.Color, w int, h int) image.Image {
	n := len(cdfs)
	if n == 0 {
		return image.NewUniform(color.Transparent)
	}

	// Find the widest range of X values:
	minX, maxX := findRange(cdfs, p)

	// Sample density for all cdfs, look for the max Y value.
	maxY := 0.0
	samples := make([][]float64, n)
	for i, cdf := range cdfs {
		samples[i] = UniformDensitySamples(cdf, minX, maxX, w)
		for _, y := range samples[i] {
			if y > maxY {
				maxY = y
			}
		}
	}

	// For each CDF: Generate image, layer on top of existing images.
	var img *image.RGBA
	bounds := image.Rect(0, 0, w, h)
	h -= 1
	m := float64(h) / maxY
	for i, sample := range samples {
		newImg := image.NewRGBA(bounds)
		var lastY int
		for x, val := range sample {
			newY := h - int(val*m)
			var y1, y2 int
			if x == 0 {
				y1, y2 = newY, newY
			} else if newY > lastY {
				y1, y2 = lastY, newY
			} else {
				y1, y2 = newY, lastY
			}
			for y := y1; y <= y2; y += 1 {
				newImg.Set(x, y, lineColors[i])
			}
			for y := y2 + 1; y <= h; y += 1 {
				newImg.Set(x, y, fillColors[i])
			}
			lastY = newY
		}
		if img == nil {
			img = newImg
		} else {
			draw.Draw(img, bounds, newImg, image.ZP, draw.Over)
		}
	}
	return img
}
