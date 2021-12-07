package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestOutputEfficientCost(t *testing.T) {
	testInputs := []int{16,1,2,0,4,2,7,1,2,14}
	minCost := outputEfficientCost(16, testInputs)
	require.Equal(t, 168, minCost)
}
