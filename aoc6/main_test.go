package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSimulate(t *testing.T) {
	testInput := []int{3,4,3,1,2}
	numFish := simulateLanternFish(testInput, 18)
	require.Equal(t, uint64(26), numFish)

	numFish = simulateLanternFish(testInput, 80)
	require.Equal(t, uint64(5934), numFish)
}
