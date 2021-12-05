package data

type Line struct {
	X1, Y1, X2, Y2 int
}

type LineDirection = string

const (
	Horizontal LineDirection = "horizontal"
	Vertical   LineDirection = "vertical"
	Diagonal   LineDirection = "diagonal"
)

type LinePoints map[int]map[int]int
