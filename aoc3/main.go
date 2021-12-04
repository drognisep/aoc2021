package main

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

const (
	fScreen uint64 = 0xffffffffffffffff
)

type DigitStats struct {
	Zeros int
	Ones  int
}

func (s DigitStats) mostCommonDigit() int {
	if s.Zeros > s.Ones {
		return 0
	}
	return 1
}

func (s DigitStats) mostCommonTieOne() int {
	if s.Zeros == s.Ones {
		return 1
	}
	return s.mostCommonDigit()
}

func (s DigitStats) leastCommonDigit() int {
	if s.mostCommonDigit() == 0 {
		return 1
	}
	return 0
}

func (s DigitStats) leastCommonTieZero() int {
	if s.Zeros == s.Ones {
		return 0
	}
	return s.leastCommonDigit()
}

func main() {
	rawInput := strings.Split(input, "\n")
	digits := len(rawInput[0])
	inputs := parseRawInputs(rawInput)
	gamma, epsilon := getGammaEpsilon(inputs, digits)
	// Gamma: 805, Epsilon: 3290, Power consumption: 2648450
	fmt.Printf("Gamma: %d, Epsilon: %d, Power consumption: %d\n", gamma, epsilon, gamma*epsilon)
	o2 := filterMostCommon(inputs, digits)
	co2 := filterLeastCommon(inputs, digits)
	fmt.Printf("Oxygen generation: %d, Co2 scrubbing: %d, Life support: %d\n", o2, co2, o2 * co2)
}

func getGammaEpsilon(inputs []uint64, digits int) (gamma uint64, epsilon uint64) {
	dataMap := getDigitStats(inputs, digits)
	for offset, stats := range dataMap {
		gamma += offset * uint64(stats.mostCommonDigit())
	}

	epsilon = gamma ^ fScreen
	epsilon = epsilon & (fScreen >> (64 - digits))

	return gamma, epsilon
}

func getDigitStats(inputs []uint64, digits int) map[uint64]*DigitStats {
	dataMap := map[uint64]*DigitStats{}

	for _, num := range inputs {
		for offset := digits - 1; offset >= 0; offset-- {
			offsetVal := uint64(1) << offset
			data, ok := dataMap[offsetVal]
			if !ok {
				data = &DigitStats{}
				dataMap[offsetVal] = data
			}
			if num&offsetVal > 0 {
				data.Ones++
			} else {
				data.Zeros++
			}
		}
	}
	return dataMap
}

func parseRawInputs(inputs []string) []uint64 {
	list := make([]uint64, len(inputs))
	for i, in := range inputs {
		num, err := strconv.ParseUint(in, 2, 64)
		if err != nil {
			log.Panicf("Input '%s' cannot be parsed as a binary number: %v\n", in, err)
		}
		list[i] = num
	}
	return list
}

func filterMostCommon(inputs []uint64, digits int) uint64 {
	dataMap := getDigitStats(inputs, digits)
	digitScreen := uint64(1) << (digits - 1)
	mostCommon := dataMap[digitScreen].mostCommonTieOne()

	var remaining []uint64
	for _, in := range inputs {
		if mostCommon == 1 {
			if in & digitScreen > 0 {
				remaining = append(remaining, in)
			}
		} else if in & digitScreen == 0 {
			remaining = append(remaining, in)
		}
	}
	if len(remaining) > 1 {
		return filterMostCommon(remaining, digits - 1)
	}
	return remaining[0]
}

func filterLeastCommon(inputs []uint64, digits int) uint64 {
	dataMap := getDigitStats(inputs, digits)
	digitScreen := uint64(1) << (digits - 1)
	leastCommon := dataMap[digitScreen].leastCommonTieZero()

	var remaining []uint64
	for _, in := range inputs {
		if leastCommon == 1 {
			if in & digitScreen > 0 {
				remaining = append(remaining, in)
			}
		} else if in & digitScreen == 0 {
			remaining = append(remaining, in)
		}
	}
	if len(remaining) > 1 {
		return filterLeastCommon(remaining, digits - 1)
	}
	return remaining[0]
}
