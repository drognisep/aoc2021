package process

import (
	"github.com/drognisep/aoc2021/aoc5/data"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestExtractPoints(t *testing.T) {
	testInput := []*data.Line{
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

	lineCh := make(chan *data.Line)
	result := make(chan data.LinePoints)
	ExtractPoints(lineCh, result)

	for _, line := range testInput {
		lineCh <- line
	}
	close(lineCh)
	lp := <-result

	require.Equal(t, 5, lp.OverlappingPoints())
}
