package strip

import (
	"fmt"
	"log"
	"strconv"

	"periph.io/x/conn/physic"
	"periph.io/x/conn/spi"
	"periph.io/x/conn/spi/spireg"
	"periph.io/x/devices/apa102"
	"periph.io/x/host"
)

type RGB struct {
	red   float64
	green float64
	blue  float64
}

type Strip struct {
	NumPixles int
	strip     *apa102.Dev
	buffer    []byte
}

func NewStrip(numPixels int, mhz int64) (Strip, error) {
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	s1, err := spireg.Open("/dev/spidev0.0")
	if err != nil {
		panic(err)
	}
	defer s1.Close()
	dd := physic.MegaHertz
	dd.Set(strconv.FormatInt(*mhz, 10) + "MHz")

	if err := s1.LimitSpeed(dd); err != nil {
		fmt.Println(err)
	}

	if p, ok := s1.(spi.Pins); ok {
		log.Printf("Using pins CLK: %s  MOSI: %s  MISO: %s", p.CLK(), p.MOSI(), p.MISO())
	}

	opts := apa102.PassThruOpts
	opts.NumPixels = 144
	opts.Intensity = 50
	a, err := apa102.New(s1, &opts)
	defer a.Halt()

	if err != nil {
		panic(err)
	}
	return Strip{
		NumPixles: numPixels,
		strip:     a,
		buffer:    []byte{},
	}, nil
}
