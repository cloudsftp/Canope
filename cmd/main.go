package main

import (
	"image/png"
	"log"
	"math"
	"os"

	canope "github.com/cloudsftp/FractalCanope"
)

func main() {
	size := 1024
	split := 3
	levels := 5
	length := 900.
	ratio := 2. / 4
	bredth := 3 * math.Pi / 4
	img := canope.DrawCanope(size, split, levels, length, ratio, bredth)

	pic, err := os.Create("canope.png")
	if err != nil {
		log.Fatalf("problem creating file: %v", err)
	}

	png.Encode(pic, img)
}
