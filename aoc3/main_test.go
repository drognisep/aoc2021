package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	testInputs = []string{
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
)

func TestGetGammaEpsilon(t *testing.T) {
	numInputs := parseRawInputs(testInputs)
	assert.Equal(t, len(testInputs), len(numInputs))
	gamma, epsilon := getGammaEpsilon(numInputs, 5)
	assert.Equal(t, uint64(22), gamma)
	assert.Equal(t, uint64(9), epsilon)
	assert.Equal(t, uint64(198), epsilon*gamma)
}

func TestLifeSupport(t *testing.T) {
	numInputs := parseRawInputs(testInputs)
	o2 := filterMostCommon(numInputs, 5)
	assert.Equal(t, uint64(23), o2)

	co2 := filterLeastCommon(numInputs, 5)
	assert.Equal(t, uint64(10), co2)

	lifeSupport := o2 * co2
	assert.Equal(t, uint64(230), lifeSupport)
}
