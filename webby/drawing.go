package main

import (
	"canvas"
	"image"
)

func DrawWeb(lines []line, canvasWidth, numberRadii int) image.Image {
	var spiralCount int
	c := canvas.CreateNewCanvas(canvasWidth, canvasWidth)

	c.SetFillColor(canvas.MakeColor(0, 0, 0))
	c.ClearRect(0, 0, canvasWidth, canvasWidth)
	c.Fill()

	c.MoveTo(lines[0].startPoint.x, lines[0].startPoint.y)
	c.SetLineWidth(5.0)
	c.SetStrokeColor(canvas.MakeColor(255, 255, 255))

	length := len(lines)

	for i := 0; i < length; i++ {
		if lines[i].lineType == "radii" {
			c.MoveTo(500.0, 500.0)
			x := lines[i].endPoint.x
			y := lines[i].endPoint.y
			c.LineTo(x, y)
		} else if lines[i].lineType == "frame" {
			x := lines[i].startPoint.x
			y := lines[i].startPoint.y
			c.LineTo(x, y)

			if i == (numberRadii*2)-1 {
				x := lines[0].endPoint.x
				y := lines[0].endPoint.y
				c.LineTo(x, y)
			}

		} else {
			spiralCount += 1
			if i == numberRadii*2 {
				x1 := lines[i].startPoint.x
				y1 := lines[i].startPoint.y
				c.MoveTo(x1, y1)
			}
			x1 := lines[i].endPoint.x
			y1 := lines[i].endPoint.y
			c.LineTo(x1, y1)

		}
	}
	c.Stroke()
	c.FillStroke()
	return c.GetImage()
}
