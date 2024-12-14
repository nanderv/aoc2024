package common

import (
	"github.com/llgcode/draw2d/draw2dimg"
	"image"
	"image/color"
)

func VisualisePos(ps []Pos, file string) {
	mX := 0
	mY := 0
	for _, p := range ps {
		if p.X > mX {
			mX = p.X
		}
		if p.Y > mY {
			mY = p.Y
		}
	}
	// Initialize the graphic context on an RGBA image
	dest := image.NewRGBA(image.Rect(0, 0, mX, mY))
	gc := draw2dimg.NewGraphicContext(dest)

	// Set some properties
	gc.SetFillColor(color.RGBA{0x44, 0xff, 0x44, 0xff})
	gc.SetStrokeColor(color.RGBA{0x44, 0x44, 0x44, 0xff})
	gc.SetLineWidth(1)
	for _, r := range ps {
		gc.BeginPath()
		gc.MoveTo(float64(r.X), float64(r.Y))
		gc.LineTo(float64(r.X+1), float64(r.Y+1))
		gc.Close()
		gc.FillStroke()
	}
	// Save to file
	err := draw2dimg.SaveToPngFile(file, dest)
	if err != nil {
		panic(err)
	}
}
