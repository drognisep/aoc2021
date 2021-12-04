package main

import (
	"github.com/drognisep/aoc2021/aoc1B/processes"
	"github.com/stretchr/testify/require"
	"testing"
)

var (
	testInputs = []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
)

func TestLogic(t *testing.T) {
	groupInput := make(chan int)
	groupOutput := make(chan int)
	resultOutput := make(chan int)

	processes.Grouper(groupInput, groupOutput)
	processes.Comparer(groupOutput, resultOutput)

	for _, i := range testInputs {
		groupInput <- i
	}
	close(groupInput)
	result := <-resultOutput
	t.Log("There were", result, "increases")
	require.Equal(t, 5, result)
}
