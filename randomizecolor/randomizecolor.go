package randomizecolor

import (
	"math/rand"
	"time"

	"github.com/wbergg/go-apa102-demo/strip"
)

func RandomizeColor() strip.RGB {
	rand.Seed(time.Now().UTC().UnixNano())

	tuple := genRGBTuple(2)
	//Return struct using struct from strip package
	return strip.RGB{
		Red:   float64(tuple[0] * 127),
		Green: float64(tuple[1] * 127),
		Blue:  float64(tuple[3] * 127),
	}
}

func genRGBTuple(max int) []int {
	t := []int{
		rand.Intn(max),
		rand.Intn(max),
		rand.Intn(max),
	}
	if t[0] == 0 && t[1] == 0 && t[2] == 0 {
		return genRGBTuple(max)
	}
	return t
}
