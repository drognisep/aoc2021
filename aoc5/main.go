package main

import (
	"bufio"
	_ "embed"
	"github.com/drognisep/aoc2021/aoc5/data"
	"github.com/drognisep/aoc2021/aoc5/process"
	"log"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	lines := parseInput(input)

	lineCh := make(chan *data.Line)
	result := make(chan data.LinePoints)
	process.ExtractPoints(lineCh, result)

	for _, line := range lines {
		lineCh <- line
	}
	close(lineCh)
	lp := <-result

	log.Printf("There are %d points overlapping\n", lp.OverlappingPoints())
}

func parseInput(input string) []*data.Line {
	scanner := bufio.NewScanner(strings.NewReader(input))
	var lines []*data.Line
	for scanner.Scan() {
		coords := strings.Split(scanner.Text(), " -> ")
		if len(coords) != 2 {
			log.Panicf("Unexpected coords length: %d", len(coords))
		}
		x1, y1 := parseCoords(coords[0])
		x2, y2 := parseCoords(coords[1])
		lines = append(lines, &data.Line{X1: x1, Y1: y1, X2: x2, Y2: y2})
	}
	return lines
}

func parseCoords(s string) (int, int) {
	coords := strings.Split(s, ",")
	if len(coords) != 2 {
		log.Panicf("Unexpected coord length: %d", len(coords))
	}
	x, err := strconv.Atoi(coords[0])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(coords[1])
	if err != nil {
		panic(err)
	}
	return x, y
}
