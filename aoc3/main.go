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

func (s DigitStats) leastCommonDigit() int {
	if s.mostCommonDigit() == 0 {
		return 1
	}
	return 0
}

func main() {
	inputs := strings.Split(input, "\n")
	gamma, epsilon := getGammaEpsilon(inputs)
	fmt.Println("Power consumption is:", gamma*epsilon)
}

func getGammaEpsilon(inputs []string) (gamma uint64, epsilon uint64) {
	dataMap := map[uint64]*DigitStats{}
	digits := len(inputs[0])

	for _, in := range inputs {
		num, err := strconv.ParseUint(in, 2, 64)
		if err != nil {
			log.Panicf("Input '%s' cannot be parsed as a binary number: %v\n", in, err)
		}

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
	for offset, stats := range dataMap {
		gamma += offset * uint64(stats.mostCommonDigit())
	}

	epsilon = gamma ^ fScreen
	epsilon = epsilon & (fScreen >> (64 - digits))

	return gamma, epsilon
}
