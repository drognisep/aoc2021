package main

import (
	"github.com/drognisep/aoc2021/aoc5/data"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParse(t *testing.T) {
	expected := []*data.Line{
		{0, 9, 5, 9},
		{8, 0, 0, 8},
		{9, 4, 3, 4},
		{2, 2, 2, 1},
		{7, 0, 7, 4},
		{6, 4, 2, 0},
		{0, 9, 2, 9},
		{3, 4, 1, 4},
		{0, 0, 8, 8},
		{5, 5, 8, 2},
	}
	actual :=
		`0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`

	lines := parseInput(actual)
	assert.Equal(t, len(expected), len(lines))

	for i, line := range lines {
		assert.Equal(t, line.X1, expected[i].X1)
		assert.Equal(t, line.Y1, expected[i].Y1)
		assert.Equal(t, line.X2, expected[i].X2)
		assert.Equal(t, line.Y2, expected[i].Y2)
	}
}
