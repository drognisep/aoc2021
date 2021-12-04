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

func main() {
	var commands []string
	for _, cmd := range strings.Split(input, "\n") {
		commands = append(commands, cmd)
	}
	result := interpret(commands)
	log.Println(result.String())
}

type Result struct {
	depth    int
	distance int
	aim      int
}

func (r Result) String() string {
	return fmt.Sprintf("Went %d at a depth of %d. The product is %d.", r.distance, r.depth, r.depth*r.distance)
}

func interpret(commands []string) Result {
	var depth, distance, aim int
	for _, cmd := range commands {
		segments := strings.SplitN(cmd, " ", 2)
		direction, amount := segments[0], segments[1]
		amountInt, err := strconv.Atoi(amount)
		if err != nil {
			log.Panicln("Failed to convert", amount, "to an integer", err)
		}
		switch direction {
		case "forward":
			distance += amountInt
			depth += aim * amountInt
		case "down":
			//depth += amountInt
			aim += amountInt
		case "up":
			//depth -= amountInt
			aim -= amountInt
		}
	}
	return Result{distance: distance, depth: depth}
}
