package main

import (
	"embed"
	"github.com/drognisep/aoc2021/aoc4/simulator"
)

//go:embed input.txt
var input embed.FS

func main() {
	data, err := input.Open("input.txt")
	if err != nil {
		panic(err)
	}
	simulator.SimulateCowardlyGameData(data)
}
