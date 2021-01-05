package main

import (
	"flag"
	"fmt"
	"math/rand"

	"./strip"
)

var mhz = flag.Int64("megahertz", 6, "what mhz to clock SPI at")

func main() {

	numPixels := 144
	pixels := make([]strip.RGB, numPixels)

	s, err := strip.NewStrip(numPixels, *mhz)
	if err != nil {
		panic(err)
	}

	var b1 = []byte{}
	for {

		// Build an array of RBG values
		c := RandomizeColor(*pixels)
		fmt.Println(c)

		for i := 0; i < opts.NumPixels; i++ {
			b1 = append(b1, []byte{
				Clamp255(c.red * 255),
				Clamp255(c.green * 255),
				Clamp255(c.blue * 255),
				byte(i),
			}...)
			//a.Write(b1)
			//fmt.Println(b1)
		}

		a.Write(b1)
		b1 = b1[:0]

	}
}

func Render() {
	fmt.Println("paint")
}
func RandomizeColor(RBG) RGB {
	r := float64(rand.Intn(255))
	g := float64(rand.Intn(255))
	b := float64(rand.Intn(255))
	return strip.RGB{r, g, b}
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
