package main

import (
	"flag"
	"time"

	"github.com/wbergg/go-apa102-demo/randomizecolor"
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
	defer s.Close()

	for {
		// Get random RGB values
		c := randomizecolor.RandomizeColor()

		for i, _ := range pixels {
			pixels[i] = c
			time.Sleep(time.Millisecond * 10)
		}
		//for i, _ := range pixels {
		//	pixels[i] = strip.RGB{
		//		Red:   0.0,
		//		Green: 0.0,
		//		Blue:  255.0,
		//	}
		//}
		s.Render(pixels)
	}
}
