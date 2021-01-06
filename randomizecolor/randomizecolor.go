package randomizecolor

import (
	"math/rand"

	"github.com/wbergg/go-apa102-demo/strip"
)

func RandomizeColor() strip.RGB {
	// Randomize a 0-255 integer
	r := float64(rand.Intn(255))
	// Round off the integer to either 0, 128 or 255
	if r <= 85 {
		r = 0
	}
	if r >= 86 || r >= 170 {
		r = 128
	}
	if r >= 171 {
		r = 255
	}

	// Randomize a 0-255 integer
	g := float64(rand.Intn(255))
	// Round off the integer to either 0, 128 or 255
	if g <= 85 {
		g = 0
	}
	if g >= 86 || g >= 170 {
		g = 128
	}
	if g >= 171 {
		g = 255
	}

	// Randomize a 0-255 integer
	b := float64(rand.Intn(255))
	// Round off the integer to either 0, 128 or 255
	if b <= 85 {
		b = 0
	}
	if b >= 86 || b >= 170 {
		b = 128
	}
	if b >= 171 {
		b = 255
	}

	//Return struct using struct from strip package
	return strip.RGB{
		Red:   r,
		Green: g,
		Blue:  b,
	}
}
