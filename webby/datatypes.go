package main

type OrderedPair struct {
	x float64
	y float64
}

type line struct {
	startPoint OrderedPair
	endPoint   OrderedPair
	angle      float64
	lineType   string
}

type Web struct {
	lines *[]line
	width int
}
