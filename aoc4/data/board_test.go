package data

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewBingoBoard(t *testing.T) {
	dataCases := map[string]struct {
		creationData   []int
		shouldPanic    bool
		markedNumbers  []int
		shouldWin      bool
		shouldWinAfter int
		winningScore   int
	}{
		"Win on row after all top row numbers are called": {
			creationData:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 16, 17, 18, 19, 10, 26, 27, 28, 29, 20, 36, 37, 38, 39, 30},
			markedNumbers:  []int{1, 2, 3, 4, 5},
			shouldWin:      true,
			shouldWinAfter: 4,
			winningScore:   420,
		},
		"Win on column after all left column numbers are called": {
			creationData:   []int{1, 16, 17, 18, 19, 2, 26, 27, 28, 29, 3, 36, 37, 38, 39, 4, 46, 47, 48, 49, 5, 56, 57, 58, 59},
			markedNumbers:  []int{1, 2, 3, 4, 5},
			shouldWin:      true,
			shouldWinAfter: 4,
			winningScore:   750,
		},
		"Win on column after all left column numbers are called including a number that doesn't exist": {
			creationData:   []int{1, 16, 17, 18, 19, 2, 26, 27, 28, 29, 3, 36, 37, 38, 39, 4, 46, 47, 48, 49, 5, 56, 57, 58, 59},
			markedNumbers:  []int{1, 2, 3, 4, 100, 5},
			shouldWin:      true,
			shouldWinAfter: 5,
			winningScore:   750,
		},
		"Panic with <25 data elements": {
			creationData: []int{1, 2, 3},
			shouldPanic:  true,
		},
		"Panic with nil data elements": {
			creationData: nil,
			shouldPanic:  true,
		},
	}

	for name, tc := range dataCases {
		t.Run(name, func(t *testing.T) {
			var board *BingoBoard
			if tc.shouldPanic {
				require.Panics(t, func() {
					board = NewBingoBoard(tc.creationData)
				})
				return
			}

			require.NotPanics(t, func() {
				board = NewBingoBoard(tc.creationData)
			})

			assert.False(t, board.HasWon())
			assert.Equal(t, 0, board.Score())

			for i, num := range tc.markedNumbers {
				board.Mark(num)
				if tc.shouldWin && tc.shouldWinAfter == i {
					require.True(t, board.HasWon())
					require.Equal(t, tc.winningScore, board.Score())
				}
			}
			assert.Equal(t, tc.shouldWin, board.HasWon())
		})
	}
}
