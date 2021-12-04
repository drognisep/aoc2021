package data

import "log"

const (
	BoardRows    = 5
	BoardColumns = 5
	BoardSize    = BoardColumns * BoardRows
)

type InputNumber = int
type BoardIndex = int

type BingoBoard struct {
	data     [BoardSize]bool
	indexMap map[InputNumber]BoardIndex
}

func NewBingoBoard(data []int) *BingoBoard {
	if len(data) != BoardSize {
		log.Panicf("Data length is %d, but should be %d\n", len(data), BoardSize)
	}

	board := BingoBoard{
		indexMap: map[InputNumber]BoardIndex{},
	}
	for i, d := range data {
		board.indexMap[d] = i
	}
	return &board
}

func (b *BingoBoard) Mark(num InputNumber) {
	idx, ok := b.indexMap[num]
	if !ok {
		return
	}
	b.data[idx] = true
}

func (b *BingoBoard) HasWon() bool {
	return b.winningColumn() || b.winningRow()
}

func (b *BingoBoard) winningColumn() bool {
outer:
	for col := 0; col < BoardColumns; col++ {
		for i := 0; i < BoardSize; i += BoardColumns {
			if !b.data[i+col] {
				continue outer
			}
		}
		return true
	}
	return false
}

func (b *BingoBoard) winningRow() bool {
outer:
	for row := 0; row < BoardSize-BoardColumns; row += BoardColumns {
		for i := 0; i < BoardColumns; i++ {
			if !b.data[i+row] {
				continue outer
			}
		}
		return true
	}
	return false
}

func (b *BingoBoard) Score() int {
	if !b.HasWon() {
		return 0
	}
	var score int
	for number, index := range b.indexMap {
		if !b.data[index] {
			score += number
		}
	}
	return score
}
