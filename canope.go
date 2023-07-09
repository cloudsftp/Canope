package canope

import (
	"image"
	"image/color"
	"math"

	"github.com/llgcode/draw2d/draw2dimg"
)

func DrawCanope(res, split int, length, ratio, bredth float64) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, res, res))
	gc := draw2dimg.NewGraphicContext(img)

	gc.SetStrokeColor(color.White)
	gc.SetLineWidth(3)

	start := (float64(res) - length) / 2.
	middle := float64(res) / 2.
	drawRecursive(
		gc,
		middle, start,
		length, 0, ratio,
		bredth, split, 0, 1,
	)

	return img
}

func drawRecursive(
	gc *draw2dimg.GraphicContext,
	startX, startY,
	length, angle, ratio, bredth float64,
	split, level, limit int,
) {
	if level >= limit {
		return
	}
	level += 1

	drawLine(gc, startX, startY, length, angle)
}

func drawLine(
	gc *draw2dimg.GraphicContext,
	startX, startY,
	length, angle float64,
) {
	targetX, targetY := calcTarget(startX, startY, length, angle)

	gc.BeginPath()
	gc.MoveTo(startX, startY)
	gc.LineTo(targetX, targetY)
	gc.Close()
	gc.FillStroke()
}

func calcTarget(x, y, length, angle float64) (targetX, targetY float64) {
	targetX = x + math.Sin(angle)*length
	targetY = y + math.Cos(angle)*length
	return
}
