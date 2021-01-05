package main

import (
	"flag"
	"fmt"

	"github.com/wbergg/go-apa102-demo/strip"
)

var mhz = flag.Int64("megahertz", 6, "what mhz to clock SPI at")

func main() {

	numPixels := 144
	var Intensity uint8 = 50
	pixels := make([]strip.RGB, numPixels, Intensity)

	s, err := strip.NewStrip(numPixels, Intensity, *mhz)
	if err != nil {
		panic(err)
	}

	var b1 = []byte{}
	for {

		// Build an array of RGB values
		c := strip.RandomizeColor()
		fmt.Println(c)

		for i := 0; i < s.NumPixles; i++ {
			b1 = append(b1, []byte{
				Clamp255(c.red * 255),
				Clamp255(c.green * 255),
				Clamp255(c.blue * 255),
				byte(i),
			}...)
			//a.Write(b1)
			//fmt.Println(b1)
		}

		s.strip.Write(b1)
		b1 = b1[:0]

	}
}

func Render() {
	fmt.Println("paint")
}

func Clamp255(v float64) byte {
	if v < 0 {
		return 0
	}
	if v > 255 {
		return 255
	}

	return byte(v)
}
