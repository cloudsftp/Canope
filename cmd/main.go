package main

import (
	"image/png"
	"log"
	"math"
	"os"

	canope "github.com/cloudsftp/FractalCanope"
)

func main() {
	img := canope.DrawCanope(1024, 2, 900, 2/3, math.Pi)

	pic, err := os.Create("canope.png")
	if err != nil {
		log.Fatalf("problem creating file: %v", err)
	}

	png.Encode(pic, img)
}
