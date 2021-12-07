package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	inputs := parseInput(input)
	outputEfficientCost(maxDistance(inputs), inputs)
}

func maxDistance(inputs []int) int {
	max := 0
	for _, i := range inputs {
		if i > max {
			max = i
		}
	}
	return max
}

func outputEfficientCost(maxDist int, inputs []int) int {
	bestBase := 0
	minCost := 0

	for base := 1; base <= (maxDist/2); base++ {
		cost := 0
		for _, val := range inputs {
			distance := int(math.Abs(float64(base - val)))
			for step := 1; step <= distance; step++ {
				cost += step
			}
		}
		if minCost == 0 || cost < minCost {
			minCost = cost
			bestBase = base
		}
	}
	fmt.Println("The most efficient position is", bestBase, "with a cost of", minCost)
	return minCost
}

func parseInput(input string) []int {
	segments := strings.Split(input, ",")
	var inputs = make([]int, len(segments))
	for i, seg := range segments {
		in, err := strconv.Atoi(seg)
		panicErr(err)
		inputs[i] = in
	}
	return inputs
}

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}
