package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetGammaEpsilon(t *testing.T) {
	inputs := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}

	gamma, epsilon := getGammaEpsilon(inputs)
	assert.Equal(t, uint64(22), gamma)
	assert.Equal(t, uint64(9), epsilon)
	assert.Equal(t, uint64(198), epsilon*gamma)
}
