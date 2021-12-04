package simulator

import (
	"bufio"
	"fmt"
	"github.com/drognisep/aoc2021/aoc4/data"
	"io"
	"log"
	"strconv"
	"strings"
)

func SimulateGameData(source io.Reader) (winningScore int) {
	var markedNumbers []int
	scanner := bufio.NewScanner(source)

	if !scanner.Scan() {
		panicNoData()
	}
	markedNumbers = parseMarkLine(scanner.Text())
	if !scanner.Scan() {
		panicNoData()
	}

	boards := parseGameBoards(scanner)
	for _, mark := range markedNumbers {
		for i, board := range boards {
			board.Mark(mark)
			if board.HasWon() {
				winningScore = board.Score() * mark
				fmt.Printf("Board %d has won with a score of %d\n", i, winningScore)
				return winningScore
			}
		}
	}
	fmt.Println("None of the boards have won!")
	return 0
}

func parseGameBoards(scanner *bufio.Scanner) (boards []*data.BingoBoard) {
	for {
		var boardNumbers []int
		for i := 0; i < data.BoardRows; i++ {
			if !scanner.Scan() {
				panicNoData()
			}
			line := scanner.Text()
			numStrs := strings.Split(line, " ")
			for _, s := range numStrs {
				if s == "" || s == " " {
					continue
				}
				num, err := strconv.Atoi(strings.TrimSpace(s))
				if err != nil {
					log.Panicln("Invalid numeric string input", s, "to parse game state")
				}
				boardNumbers = append(boardNumbers, num)
			}
		}

		boards = append(boards, data.NewBingoBoard(boardNumbers))

		if !scanner.Scan() {
			return boards
		}
	}
}

func parseMarkLine(line string) []int {
	numStrs := strings.Split(line, ",")
	markNums := make([]int, len(numStrs))

	for i, s := range numStrs {
		num, err := strconv.Atoi(s)
		if err != nil {
			log.Panicf("Unable to parse '%s' into int: %v\n", s, err)
		}
		markNums[i] = num
	}
	return markNums
}

func panicNoData() {
	log.Panicln("Unexpected end of data")
}
