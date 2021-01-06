package main

import (
	"flag"
	"fmt"
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

	index := 0
	direction := false
	// Get random RGB values
	//c := randomizecolor.RandomizeColor()

	for {

		c := randomizecolor.RandomizeColor()
		fmt.Println(c)

		if index >= numPixels {
			direction = false
		}
		if index <= 0 {
			direction = true
		}

		if direction {
			index++
		} else {
			index--
		}

		for i, _ := range pixels {

			if i == index {
				pixels[i] = c
				s.Render(pixels)
				time.Sleep(time.Millisecond * 10)
			} else {
				pixels[i] = c
				s.Render(pixels)
				time.Sleep(time.Millisecond * 10)
			}
		}

		//s.Render(pixels)
	}
}

//// Different patterns, split these to separate packages later
// Light all, fixed color
//func LitAll() {
//for i, _ := range pixels {
//	pixels[i] = strip.RGB{
//		Red:   0.0,
//		Green: 0.0,
//		Blue:  255.0,
//	}
//}
//s.Render(pixels)
//}
