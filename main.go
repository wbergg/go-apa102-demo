package main

import (
	"flag"

	"github.com/wbergg/go-apa102-demo/strip"
)

var mhz = flag.Int64("megahertz", 6, "what mhz to clock SPI at")

func main() {

	numPixels := 144
	var Intensity uint8 = 50
	pixels := make([]strip.RGB, numPixels)

	s, err := strip.NewStrip(numPixels, Intensity, *mhz)
	if err != nil {
		panic(err)
	}

	for {
		// Random RGB values
		//c := strip.RandomizeColor()
		for i, _ := range pixels {
			pixels[i] = strip.RGB{
				Red:   1.0,
				Green: 1.0,
				Blue:  1.0,
			}
		}
		s.Render(pixels)
	}
}
