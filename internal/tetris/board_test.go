package tetris

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBoard(t *testing.T) {
	board := NewBoard(5, 10)
	assert.Equal(t, 5, board.width)
	assert.Equal(t, 10, board.height)
	assert.False(t, board.BlockExists(0, 0))
	assert.False(t, board.BlockExists(4, 0))
	assert.False(t, board.BlockExists(0, 9))
	assert.False(t, board.BlockExists(4, 9))
	assert.Equal(t, 5, len(board.blocks))
	assert.Equal(t, 10, len(board.blocks[0]))
}

func TestPutBlock(t *testing.T) {
	board := NewBoard(5, 10)
	assert.False(t, board.Put(0, 10, 1, 1))
	assert.False(t, board.BlockExists(0, 0))
	assert.True(t, board.Put(0, 0, 1, 1))
	assert.True(t, board.BlockExists(0, 0))
	assert.True(t, board.Put(0, 0, 1, 1))
	assert.True(t, board.BlockExists(0, 1))
	assert.True(t, board.Put(0, 2, 1, 1))
	assert.True(t, board.BlockExists(0, 2))
	assert.True(t, board.Put(0, 4, 1, 1))
	assert.False(t, board.BlockExists(0, 4))
	assert.True(t, board.BlockExists(0, 3))
	assert.True(t, board.Put(0, 9, 1, 1))
	assert.False(t, board.BlockExists(0, 9))
	assert.True(t, board.BlockExists(0, 4))
	// 置き場所がボードの縦サイズを超える
	assert.False(t, board.Put(0, 3, 1, 7))
}
