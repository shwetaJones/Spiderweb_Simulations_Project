package main

import (
	"math"
	"math/rand"
)

// Generates all the lines that make up the spider web, this is made with
// Input: initial angle, canvas wdith, and the number of radii
// Output: all the lines that make up the web
func CreateWebRandom(initialAngle, canvasWidth float64, numberRadii int, alpha float64, minDistance float64, initialDistance float64) []line {
	// determines and collects the radii points
	radii := CreateRadiiRandom(initialAngle, numberRadii, canvasWidth)

	// determines and collected the frame points using radii values
	frame := CreateFrame(radii)

	// determines and collects the spiral points
	alpha += rand.Float64()
	spiral := CreateSpiralRandom(radii, minDistance, initialDistance, alpha, numberRadii, canvasWidth)
	var allLines []line

	// normalizes the radii points to be from cartesian to canvas
	normalizedRadii := make([]line, len(radii))
	for i := range radii {
		radii[i].startPoint = ChangeAxisCartToCanvas(radii[i].startPoint, canvasWidth)
		radii[i].endPoint = ChangeAxisCartToCanvas(radii[i].endPoint, canvasWidth)
		normalizedRadii[i] = radii[i]
		allLines = append(allLines, normalizedRadii[i])
	}

	// normalizes the frame points to be from cartesian to canvas
	normalizedFrame := make([]line, len(frame))
	for i := range radii {
		frame[i].startPoint = ChangeAxisCartToCanvas(frame[i].startPoint, canvasWidth)
		frame[i].endPoint = ChangeAxisCartToCanvas(frame[i].endPoint, canvasWidth)
		normalizedFrame[i] = frame[i]
		allLines = append(allLines, normalizedFrame[i])
	}

	// normalizes the spiral points to be from cartesian to canvas
	normalizedSpiral := make([]line, len(spiral))
	for i := range spiral {
		spiral[i].startPoint = ChangeAxisCartToCanvas(spiral[i].startPoint, canvasWidth)
		spiral[i].endPoint = ChangeAxisCartToCanvas(spiral[i].endPoint, canvasWidth)
		normalizedSpiral[i] = spiral[i]
		allLines = append(allLines, normalizedSpiral[i])
	}
	return allLines

}

// Generates all the lines that make up the spider web in symmeric manner
// Input: initial angle, canvas width, and the number of radii
// Output: all the lines that make up the web
func CreateWebSymmetric(initialAngle, canvasWidth float64, numberRadii int, alpha float64, minDistance float64, initialDistance float64) []line {
	// determines and collects the radii points
	radii := CreateRadiiSymmetric(initialAngle, numberRadii, canvasWidth)

	// determines and collected the frame points using radii values
	frame := CreateFrame(radii)

	// determines and collects the spiral points
	alpha += 0.5
	spiral := CreateSpiralSymmetric(radii, minDistance, initialDistance, alpha, numberRadii, canvasWidth)

	var allLines []line

	// normalizes the radii points to be from cartesian to canvas
	normalizedRadii := make([]line, len(radii))
	for i := range radii {
		radii[i].startPoint = ChangeAxisCartToCanvas(radii[i].startPoint, canvasWidth)
		radii[i].endPoint = ChangeAxisCartToCanvas(radii[i].endPoint, canvasWidth)
		normalizedRadii[i] = radii[i]
		allLines = append(allLines, normalizedRadii[i])
	}

	// normalizes the frame points to be from cartesian to canvas
	normalizedFrame := make([]line, len(frame))
	for i := range radii {
		frame[i].startPoint = ChangeAxisCartToCanvas(frame[i].startPoint, canvasWidth)
		frame[i].endPoint = ChangeAxisCartToCanvas(frame[i].endPoint, canvasWidth)
		normalizedFrame[i] = frame[i]
		allLines = append(allLines, normalizedFrame[i])
	}

	// normalizes the spiral points to be from cartesian to canvas
	normalizedSpiral := make([]line, len(spiral))
	for i := range spiral {
		spiral[i].startPoint = ChangeAxisCartToCanvas(spiral[i].startPoint, canvasWidth)
		spiral[i].endPoint = ChangeAxisCartToCanvas(spiral[i].endPoint, canvasWidth)
		normalizedSpiral[i] = spiral[i]
		allLines = append(allLines, normalizedSpiral[i])
	}
	return allLines
}

// Generates the lines that create the spirals of the web in a random fashion
// Input: the radii, minimum distance, initial distance, alpha, the number of radii, canvas width
// Output: the lines that make up the spirals of the web
func CreateSpiralRandom(radii []line, minDistance, initialDistance, alpha float64, numberRadii int, canvasWidth float64) []line {
	spiral := make([]line, 0)

	var minDistReached bool
	lineCount := 0
	radiiCount := 0
	distOffSet := 0.0
	counterClockwise := false

	var origin OrderedPair
	origin.x = 0.0
	origin.y = 0.0

	for !minDistReached {
		// initialize this line
		var newLine line

		// set start point
		if lineCount == 0 {
			// determine start of first ever spiral
			distOffSet += initialDistance
			newLine.startPoint = CalcSpiralPoint(radii[0], distOffSet)
		} else {
			// sets the start point of every spiral
			newLine.startPoint = spiral[lineCount-1].endPoint
		}

		// sets the radii for spiral endpoint
		if !counterClockwise {
			radiiCount += 1
		} else {
			radiiCount -= 1
		}
		// normalizes the radiiCount
		if radiiCount >= numberRadii {
			radiiCount = 0
		} else if radiiCount < 0 {
			radiiCount = numberRadii - 1
		}

		// the randomized value for alpha for the following radii
		alpha1 := -1.5 + rand.Float64()*3

		// sets end point
		distOffSet += alpha + alpha1

		// determines the endpoint of the spiral line
		endPoint := CalcSpiralPoint(radii[radiiCount], distOffSet)

		if CalcDistance(endPoint, radii[radiiCount].endPoint) > minDistance &&
			CalcDistance(endPoint, origin) < CalcDistance(radii[radiiCount].startPoint, radii[radiiCount].endPoint) &&
			newLine.startPoint.x >= -1*(canvasWidth/2) &&
			newLine.startPoint.x <= canvasWidth/2 &&
			endPoint.x >= -1*(canvasWidth/2) &&
			endPoint.x <= canvasWidth/2 &&
			newLine.startPoint.y >= -1*(canvasWidth/2) &&
			newLine.startPoint.y <= canvasWidth/2 &&
			endPoint.y >= -1*(canvasWidth/2) &&
			endPoint.y <= canvasWidth/2 {
			// sets the endpoint if it is a valid endpoint
			newLine.endPoint = endPoint

		} else {
			// check opposite direction since the current direction is not valid
			if !counterClockwise {
				counterClockwise = true
				radiiCount -= 2
			} else {
				counterClockwise = false
				radiiCount += 2
			}

			// resets the radiiCount
			if radiiCount >= numberRadii {
				radiiCount = 0
			} else if radiiCount < 0 {
				radiiCount = numberRadii - 1
			}

			// determines the new endpoint
			endPoint = CalcSpiralPoint(radii[radiiCount], distOffSet)

			if CalcDistance(endPoint, radii[radiiCount].endPoint) > minDistance &&
				CalcDistance(endPoint, origin) < CalcDistance(radii[radiiCount].startPoint, radii[radiiCount].endPoint) &&
				newLine.startPoint.x >= -1*(canvasWidth/2) &&
				newLine.startPoint.x <= canvasWidth/2 &&
				endPoint.x >= -1*(canvasWidth/2) &&
				endPoint.x <= canvasWidth/2 &&
				newLine.startPoint.y >= -1*(canvasWidth/2) &&
				newLine.startPoint.y <= canvasWidth/2 &&
				endPoint.y >= -1*(canvasWidth/2) &&
				endPoint.y <= canvasWidth/2 {
				// sets the endpoint if it is valid
				newLine.endPoint = endPoint
			} else {
				// at this point, both directions have met minimum distance, and the final endpoint is set as the next radii endpoint
				newLine.endPoint = radii[radiiCount].endPoint
				minDistReached = true
			}
		}

		lineCount += 1
		newLine.lineType = "spiral"
		spiral = append(spiral, newLine)
	}
	return spiral
}

// Generates the lines that create the spirals of the web in a symmetric manner
// Input: the radii, minimum distance, initial distance, alpha, the number of radii, canvas width
// Output: the lines that make up the spirals of the web
func CreateSpiralSymmetric(radii []line, minDistance, initialDistance, alpha float64, numberRadii int, canvasWidth float64) []line {
	spiral := make([]line, 0)

	var minDistReached bool
	lineCount := 0
	radiiCount := 0
	distOffSet := 0.0
	counterClockwise := false

	var origin OrderedPair
	origin.x = 0.0
	origin.y = 0.0

	for !minDistReached {
		// initialize this line
		var newLine line

		// set start point
		if lineCount == 0 {
			// determine start of first ever spiral
			distOffSet += initialDistance
			newLine.startPoint = CalcSpiralPoint(radii[0], distOffSet)
		} else {
			// sets the start point of every spiral
			newLine.startPoint = spiral[lineCount-1].endPoint
		}

		// sets the radii for spiral endpoint
		if !counterClockwise {
			radiiCount += 1
		} else {
			radiiCount -= 1
		}
		// normalizes the radiiCount
		if radiiCount >= numberRadii {
			radiiCount = 0
		} else if radiiCount < 0 {
			radiiCount = numberRadii - 1
		}

		// the randomized value for alpha for the following radii
		alpha1 := -1.5

		// sets end point
		distOffSet += alpha + alpha1

		// determines the endpoint of the spiral line
		endPoint := CalcSpiralPoint(radii[radiiCount], distOffSet)

		if CalcDistance(endPoint, radii[radiiCount].endPoint) > minDistance &&
			CalcDistance(endPoint, origin) < CalcDistance(radii[radiiCount].startPoint, radii[radiiCount].endPoint) &&
			newLine.startPoint.x >= -1*(canvasWidth/2) &&
			newLine.startPoint.x <= canvasWidth/2 &&
			endPoint.x >= -1*(canvasWidth/2) &&
			endPoint.x <= canvasWidth/2 &&
			newLine.startPoint.y >= -1*(canvasWidth/2) &&
			newLine.startPoint.y <= canvasWidth/2 &&
			endPoint.y >= -1*(canvasWidth/2) &&
			endPoint.y <= canvasWidth/2 {
			// sets the endpoint if it is a valid endpoint
			newLine.endPoint = endPoint

		} else {
			// check opposite direction since the current direction is not valid
			if !counterClockwise {
				counterClockwise = true
				radiiCount -= 2
			} else {
				counterClockwise = false
				radiiCount += 2
			}

			// resets the radiiCount
			if radiiCount >= numberRadii {
				radiiCount = 0
			} else if radiiCount < 0 {
				radiiCount = numberRadii - 1
			}

			// determines the new endpoint
			endPoint = CalcSpiralPoint(radii[radiiCount], distOffSet)

			if CalcDistance(endPoint, radii[radiiCount].endPoint) > minDistance &&
				CalcDistance(endPoint, origin) < CalcDistance(radii[radiiCount].startPoint, radii[radiiCount].endPoint) &&
				newLine.startPoint.x >= -1*(canvasWidth/2) &&
				newLine.startPoint.x <= canvasWidth/2 &&
				endPoint.x >= -1*(canvasWidth/2) &&
				endPoint.x <= canvasWidth/2 &&
				newLine.startPoint.y >= -1*(canvasWidth/2) &&
				newLine.startPoint.y <= canvasWidth/2 &&
				endPoint.y >= -1*(canvasWidth/2) &&
				endPoint.y <= canvasWidth/2 {
				// sets the endpoint if it is valid
				newLine.endPoint = endPoint
			} else {
				// at this point, both directions have met minimum distance, and the final endpoint is set as the next radii endpoint
				newLine.endPoint = radii[radiiCount].endPoint
				minDistReached = true
			}
		}

		lineCount += 1
		newLine.lineType = "spiral"
		spiral = append(spiral, newLine)
	}
	return spiral
}

// Calculates the distance between two given points
// Input: 2 points
// Output: the distance between two points
func CalcDistance(firstPoint, secondPoint OrderedPair) float64 {
	var distance float64
	deltaX := secondPoint.x - firstPoint.x
	deltaY := secondPoint.y - firstPoint.y
	distance = math.Sqrt(deltaX*deltaX + deltaY*deltaY)
	return distance
}

// Calculates the spiral point
// Input: the radius of interest and the distance offset value
// Output: the spiral point
func CalcSpiralPoint(radius line, distOffset float64) OrderedPair {
	var point OrderedPair
	point.x = (distOffset * math.Cos(radius.angle*(3.14/180)))
	point.y = ((radius.endPoint.y / radius.endPoint.x) * point.x)
	return point
}

// Calculates all the radii points
// Input: the innitial angle, the number of radii, and the width
// Output: a set of lines that compromise the radii
func CreateRadiiRandom(initialAngle float64, numberRadii int, width float64) []line {
	// determines the base points that set the limits of the radii
	x1, x2, y1, y2 := BasePointsRandom(width)

	// creates the first radii
	firstLine := FirstRadiiRandom(initialAngle, width, x1, x2, y1, y2) // CHANGED THIS
	radii := make([]line, numberRadii)
	radii[0] = firstLine
	radii[0].angle = initialAngle
	radii[0].lineType = "radii"

	// creates the rest of the radii with a degree of randomization
	equalAngle := 360.0 / float64(numberRadii)
	angleVariance := (-1 + rand.Float64()*2) * 10
	lineAngle := initialAngle + equalAngle + angleVariance
	for i := 1; i < numberRadii; i++ {
		nextAngleVariance := (-1 + rand.Float64()*2) * 10
		radii[i] = FirstRadiiRandom(lineAngle, width, x1, x2, y1, y2)
		radii[i].angle = lineAngle
		radii[i].lineType = "radii"
		lineAngle += equalAngle + nextAngleVariance
	}
	return radii
}

// Calculates all the radii points
// Input: the innitial angle, the number of radii, and the width
// Output: a set of lines that compromise the radii
func CreateRadiiSymmetric(initialAngle float64, numberRadii int, width float64) []line {
	// determines the base points that set the limits of the radii
	x1, x2, y1, y2 := BasePointsSymmetric(width)

	// creates the first radii
	firstLine := FirstRadiiSymmetric(initialAngle, width, x1, x2, y1, y2)
	radii := make([]line, numberRadii)
	radii[0] = firstLine
	radii[0].angle = initialAngle
	radii[0].lineType = "radii"

	// creates the rest of the radii
	equalAngle := 360.0 / float64(numberRadii)
	angleVariance := (-1 + 0.5*2) * 10
	lineAngle := initialAngle + equalAngle + angleVariance
	for i := 1; i < numberRadii; i++ {
		nextAngleVariance := (-1 + 0.5*2) * 10
		radii[i] = FirstRadiiSymmetric(lineAngle, width, x1, x2, y1, y2)
		radii[i].angle = lineAngle
		radii[i].lineType = "radii"
		lineAngle += equalAngle + nextAngleVariance
	}
	return radii
}

// Determines the frame points
// Input: the radii
// Output: a set of lines that compromise the frame
func CreateFrame(radii []line) []line {
	numberRadii := len(radii)
	frame := make([]line, numberRadii)
	// iterates through the radii in order to set the start and end points of the frame
	for i := 0; i < numberRadii; i++ {

		// the frame startpoint is the endpoints of it's corresponding the endpoint
		frame[i].startPoint = radii[i].endPoint

		if i+1 >= numberRadii {

			// the final frame line's endpoint is the first radii's endpoint
			frame[i].endPoint = radii[0].endPoint

		} else {

			// for all other frame's, the endpoint is the next radii's endpoint
			frame[i].endPoint = radii[i+1].endPoint

		}

		frame[i].lineType = "frame"
	}
	return frame
}

// Determines the first radii
// Input: the initial angle, the canvas width, and the basepoints
// Output: the first radii line
func FirstRadiiRandom(initialAngle float64, width float64, x1, y1, x2, y2 OrderedPair) line {
	var length float64
	var endPoint OrderedPair
	var startPoint OrderedPair
	var radius line

	// checks the quadrant in which the angle lies
	initialQuad := CheckQuadrant(initialAngle)

	// based on the quadrant, the corresponding length of the radius is determined
	if initialQuad == 1 {
		length = CalcLengthRandom(x1.x, y1.y, initialAngle)
	} else if initialQuad == 2 {
		length = CalcLengthRandom(x2.x, y1.y, initialAngle)
	} else if initialQuad == 3 {
		length = CalcLengthRandom(x2.x, y2.y, initialAngle)
	} else if initialQuad == 4 {
		length = CalcLengthRandom(x1.x, y2.y, initialAngle)
	}

	// based on the length and initial angle, the correpsonding endpoint is determined
	endPoint.x = length * math.Cos(initialAngle*(3.14/180))
	endPoint.y = length * math.Sin(initialAngle*(3.14/180))
	startPoint.x = 0.0
	startPoint.y = 0.0

	// the points are set to the radius
	radius.startPoint = startPoint
	radius.endPoint = endPoint

	return radius
}

// Determines the first radii
// Input: the initial angle, the canvas width, and the basepoints
// Output: the first radii line
func FirstRadiiSymmetric(initialAngle float64, width float64, x1, y1, x2, y2 OrderedPair) line {
	var length float64
	var endPoint OrderedPair
	var startPoint OrderedPair
	var radius line

	// the length remains uniform for all radii
	length = width / 2

	// based on the length and initial angle, the correpsonding endpoint is determined
	endPoint.x = length * math.Cos(initialAngle*(3.14/180))
	endPoint.y = length * math.Sin(initialAngle*(3.14/180))
	startPoint.x = 0.0
	startPoint.y = 0.0

	// the points are set to the radius
	radius.startPoint = startPoint
	radius.endPoint = endPoint

	return radius
}

// Determines the basepoints
// Input: the canvas width
// Output: the base points ordered pairs
func BasePointsRandom(width float64) (OrderedPair, OrderedPair, OrderedPair, OrderedPair) {
	var x1 OrderedPair
	var x2 OrderedPair
	var y1 OrderedPair
	var y2 OrderedPair

	// sets the basepoint values
	x1.x = 0.8*width + rand.Float64()*0.1*width
	x1.y = width / 2

	y1.x = width / 2
	y1.y = rand.Float64() * 0.15 * width

	x2.x = rand.Float64() * 0.15 * width
	x2.y = width / 2

	y2.x = width / 2
	y2.y = 0.8*width + rand.Float64()*0.1*width

	// converts all the canvas points to cartesian points
	x1 = ChangeAxisCanvastoCart(x1, width)
	x2 = ChangeAxisCanvastoCart(x2, width)
	y1 = ChangeAxisCanvastoCart(y1, width)
	y2 = ChangeAxisCanvastoCart(y2, width)

	return x1, y1, x2, y2
}

// Determines the basepoints, this creates set points based off the width
// Input: the canvas width
// Output: the base points ordered pairs
func BasePointsSymmetric(width float64) (OrderedPair, OrderedPair, OrderedPair, OrderedPair) {
	var x1 OrderedPair
	var x2 OrderedPair
	var y1 OrderedPair
	var y2 OrderedPair

	// sets the basepoint values
	x1.x = (3 * width) / 4.0
	x1.y = width / 2

	y1.x = width / 2
	y1.y = width / 4.0

	x2.x = width / 4.0
	x2.y = width / 2

	y2.x = width / 2
	y2.y = (3 * width) / 4.0

	// converts all the canvas points to cartesian points
	x1 = ChangeAxisCanvastoCart(x1, width)
	x2 = ChangeAxisCanvastoCart(x2, width)
	y1 = ChangeAxisCanvastoCart(y1, width)
	y2 = ChangeAxisCanvastoCart(y2, width)

	return x1, y1, x2, y2
}

// Changes coordinate points from canvas to cartesian
// Input: a canvas point and the canvas width
// Output: a cartesian point
func ChangeAxisCanvastoCart(point OrderedPair, width float64) OrderedPair {
	point.x = point.x - width/2
	point.y = (width - point.y) - width/2

	return point
}

// Changes coordinate points from cartesian to canvas
// Input: a coordinate point and the canvas width
// Output: a canvas point
func ChangeAxisCartToCanvas(point OrderedPair, width float64) OrderedPair {
	point.x = point.x + width/2
	point.y = (width + point.y) - width/2

	return point
}

// Checks the quadrant where the angle is in
// Input: a angle
// Output: the corresponding quadrant to the give angle
func CheckQuadrant(inputAngle float64) int {
	if inputAngle >= 0 && inputAngle <= 90 {
		return 1
	} else if inputAngle > 90 && inputAngle <= 180 {
		return 2
	} else if inputAngle > 180 && inputAngle <= 270 {
		return 3
	} else if inputAngle > 270 && inputAngle <= 360 {
		return 4
	} else if inputAngle > 360 {
		return CheckQuadrant(inputAngle - 360)
	} else if inputAngle < 0 {
		inputAngle = inputAngle * -1
		return CheckQuadrant(inputAngle)
	}
	return -1
}

// Calculates the length with a randomization factor
// Input: x and y coordinate and an angle
// Output: a calculated length
func CalcLengthRandom(x, y float64, angle float64) float64 {
	var a, b, length float64

	b = 0.4*y + rand.Float64()*(0.1*y)
	a = (x*x - y*y + 2*b*y) / (2 * x)
	c := a*math.Cos(angle*(3.14/180)) + b*math.Sin(angle*(3.14/180))
	length = c + math.Sqrt(c*c-(2*b*y-y*y))

	return length
}

// Calculates the length
// Input: x and y coordinate and an angle
// Output: a calculated length
func CalcLengthSymmetric(x, y float64, angle float64) float64 {
	var a, b, length float64

	b = 0.4*y + 0.5*(0.1*y)
	a = (x*x - y*y + 2*b*y) / (2 * x)
	c := a*math.Cos(angle*(3.14/180)) + b*math.Sin(angle*(3.14/180))
	length = c + math.Sqrt(c*c-(2*b*y-y*y))

	return length
}
