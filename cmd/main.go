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
	split := 2
	levels := 7
	length := 900.
	ratio := 0.4
	bredth := math.Pi * 0.3
	line := 10.
	img := canope.DrawCanope(size, split, levels, length, ratio, bredth, line)

	pic, err := os.Create("canope.png")
	if err != nil {
		log.Fatalf("problem creating file: %v", err)
	}

	png.Encode(pic, img)
}
