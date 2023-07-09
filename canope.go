package canope

import (
	"image"
	"image/color"
	"math"

	"github.com/llgcode/draw2d/draw2dimg"
)

func DrawCanope(res, split, levels int, length, ratio, bredth, line float64) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, res, res))
	gc := draw2dimg.NewGraphicContext(img)

	gc.SetStrokeColor(color.White)

	start := float64(res) - (float64(res)-length)/2
	middle := float64(res) / 2
	drawRecursive(
		gc,
		middle, start,
		length, 0, ratio,
		bredth, line, split, 0, levels,
	)

	return img
}

func drawRecursive(
	gc *draw2dimg.GraphicContext,
	startX, startY,
	length, angle, ratio, bredth, line float64,
	split, level, limit int,
) {
	if level >= limit {
		return
	}
	level += 1

	drawLine(gc, startX, startY, length*ratio, angle, line)

	startX, startY = calcTarget(startX, startY, length*ratio, angle)
	line = line * 0.7
	startAngle := angle - bredth/2
	stepAngle := bredth / float64(split-1)
	for i := 0; i < split; i++ {
		drawRecursive(
			gc,
			startX, startY,
			length*(1-ratio), startAngle+float64(i)*stepAngle, ratio,
			bredth, line, split, level, limit,
		)
	}
}

func drawLine(
	gc *draw2dimg.GraphicContext,
	startX, startY,
	length, angle, line float64,
) {
	targetX, targetY := calcTarget(startX, startY, length, angle)

	gc.SetLineWidth(line)
	gc.BeginPath()
	gc.MoveTo(startX, startY)
	gc.LineTo(targetX, targetY)
	gc.Close()
	gc.FillStroke()
}

func calcTarget(x, y, length, angle float64) (targetX, targetY float64) {
	targetX = x + math.Sin(angle)*length
	targetY = y - math.Cos(angle)*length
	return
}
