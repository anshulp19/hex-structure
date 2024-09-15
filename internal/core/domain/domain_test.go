package domain_test

import (
	"hex-structure/internal/core/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBoard(t *testing.T) {
	size := uint(10)
	bombs := uint(50)

	board := domain.NewBoard(size, bombs)
	countBombs := 0

	for row := range board {
		for col := range board[0] {
			if board[row][col] == domain.CELL_BOMB {
				countBombs++
			}
		}
	}

	assert.Equal(t, uint(len(board)), size)
	assert.Equal(t, uint(len(board[0])), size)
	assert.Equal(t, uint(countBombs), bombs)
}
