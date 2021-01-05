package strip

import (
	"fmt"
	"log"
	"math/rand"
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
	Intensity uint8
	strip     *apa102.Dev
	buffer    []byte
}

func NewStrip(numPixels int, Intensity uint8, mhz int64) (Strip, error) {
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	s1, err := spireg.Open("/dev/spidev0.0")
	if err != nil {
		panic(err)
	}
	defer s1.Close()
	dd := physic.MegaHertz
	dd.Set(strconv.FormatInt(mhz, 10) + "MHz")

	if err := s1.LimitSpeed(dd); err != nil {
		fmt.Println(err)
	}

	if p, ok := s1.(spi.Pins); ok {
		log.Printf("Using pins CLK: %s  MOSI: %s  MISO: %s", p.CLK(), p.MOSI(), p.MISO())
	}

	opts := apa102.PassThruOpts
	opts.NumPixels = numPixels
	opts.Intensity = Intensity
	a, err := apa102.New(s1, &opts)
	defer a.Halt()

	if err != nil {
		panic(err)
	}
	return Strip{
		NumPixles: numPixels,
		Intensity: Intensity,
		strip:     a,
		buffer:    []byte{},
	}, nil
}

func (s *Strip) Render(p []RGB) {
	s.buffer = s.buffer[:0]
	for i, p := range p {
		s.buffer = append(s.buffer, []byte{
			clamp255(p.red * 255),
			clamp255(p.green * 255),
			clamp255(p.blue * 255),
			byte(i),
		}...)
	}
	s.strip.Write(s.buffer)
}

func clamp255(v float64) byte {
	if v < 0 {
		return 0
	}
	if v > 255 {
		return 255
	}

	return byte(v)
}

func RandomizeColor() RGB {
	r := float64(rand.Intn(255))
	g := float64(rand.Intn(255))
	b := float64(rand.Intn(255))
	return RGB{
		red:   r,
		green: g,
		blue:  b,
	}
}
