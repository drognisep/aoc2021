package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestInterpret(t *testing.T) {
	testCommands := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}

	result := interpret(testCommands)
	require.Equal(t, 15, result.distance)
	require.Equal(t, 60, result.depth)
	require.Equal(t, "Went 15 at a depth of 60. The product is 900.", result.String())
}
