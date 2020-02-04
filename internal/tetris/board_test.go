package tetris

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBoard(t *testing.T) {
	board := NewBoard(5, 10)
	assert.Equal(t, 5, board.width)
	assert.Equal(t, 10, board.height)
	assert.False(t, board.blockExists(0,0))
	assert.False(t, board.blockExists(4,0))
	assert.False(t, board.blockExists(0,9))
	assert.False(t, board.blockExists(4,9))
	assert.Equal(t, 5, len(board.blocks))
	assert.Equal(t, 10, len(board.blocks[0]))
	assert.False(t, board.blockExists(0,0))
	assert.True(t, board.put(0,0))
	assert.True(t, board.blockExists(0,0))
	assert.False(t, board.put(0, 10))
	assert.True(t, board.put(0,0))
	assert.True(t, board.blockExists(0,1))
	assert.True(t, board.put(0,0))
	assert.True(t, board.blockExists(0,2))
	assert.True(t, board.put(0,0))
	assert.True(t, board.put(0,1))
	assert.True(t, board.blockExists(0,4))
	assert.True(t, board.put(0,1))
	assert.True(t, board.put(0,1))
	assert.True(t, board.put(0,1))
	assert.True(t, board.put(0,1))
	assert.True(t, board.put(0,1))
	// 置き場所がボードの縦サイズを超える
	assert.False(t, board.put(0,1))
}
