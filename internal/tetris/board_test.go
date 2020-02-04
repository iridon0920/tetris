package tetris

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBoard(t *testing.T) {
	board := NewBoard(5, 10)
	assert.Equal(t, 5, board.width)
	assert.Equal(t, 10, board.height)
	assert.False(t, board.blocks[0][0])
	assert.False(t, board.blocks[4][0])
	assert.False(t, board.blocks[0][9])
	assert.False(t, board.blocks[4][9])
	assert.Equal(t, 5, len(board.blocks))
	assert.Equal(t, 10, len(board.blocks[0]))
}
