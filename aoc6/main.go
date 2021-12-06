package main

import (
	_ "embed"
	"fmt"
	"github.com/drognisep/aoc2021/aoc6/store"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	segments := strings.Split(input, ",")
	fishes := make([]int, len(segments))

	for i, s := range segments {
		num, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		fishes[i] = num
	}
	numFish := simulateLanternFish(fishes, 256)
	fmt.Println("Num fish:", numFish)
}

func simulateLanternFish(input []int, days int) uint64 {
	fs := store.NewFishStore(input)

	for day := 1; day <= days; day++ {
		fs.DecrementFish()
	}
	return fs.CountFish()
}
