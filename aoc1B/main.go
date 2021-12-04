package main

import (
	"bufio"
	"embed"
	"fmt"
	"github.com/drognisep/aoc2021/aoc1B/processes"
	"log"
	"strconv"
)

//go:embed input.txt
var input embed.FS

func main() {
	file, err := input.Open("input.txt")
	if err != nil {
		panic(err)
	}

	groupInput := make(chan int)
	groupOutput := make(chan int)
	resultOutput := make(chan int)

	processes.Grouper(groupInput, groupOutput)
	processes.Comparer(groupOutput, resultOutput)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		val, err := strconv.Atoi(line)
		if err != nil {
			log.Panicln("Failed to convert", line, "to int", err)
		}
		groupInput <- val
	}
	if err := scanner.Err(); err != nil {
		log.Panicln("Scanner error", err)
	}
	close(groupInput)
	result := <-resultOutput
	fmt.Println("There were", result, "increases")
}
