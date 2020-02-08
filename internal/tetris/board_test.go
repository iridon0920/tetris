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

func TestMoveBlock(t *testing.T) {
	board := NewBoard(5, 3)
	assert.True(t, board.Insert(2))
	assert.True(t, board.BlockExists(2, 0))
	assert.Equal(t, 2, board.moveX)
	assert.Equal(t, 0, board.moveY)

	assert.True(t, board.MoveLeft())
	assert.False(t, board.BlockExists(2, 0))
	assert.True(t, board.BlockExists(1, 0))

	assert.True(t, board.MoveBottom())
	assert.False(t, board.BlockExists(1, 0))
	assert.True(t, board.BlockExists(1, 1))

	assert.True(t, board.MoveRight())
	assert.False(t, board.BlockExists(1, 1))
	assert.True(t, board.BlockExists(2, 1))

	assert.True(t, board.MoveBottom())
}

func TestInsertBlockMiss(t *testing.T) {
	board := NewBoard(5, 10)
	//範囲外
	assert.False(t, board.Insert(6))
	assert.False(t, board.Insert(-1))
	//投入場所に既にに存在する
	assert.True(t, board.Insert(0))
	assert.False(t, board.Insert(0))
	//制御中に投入する
	assert.False(t, board.Insert(2))
}
func TestMoveBlockMiss(t *testing.T) {
	board := NewBoard(3, 10)
	//範囲外
	assert.True(t, board.Insert(0))
	assert.False(t, board.MoveLeft())

	board2 := NewBoard(3, 10)
	//範囲外
	assert.True(t, board2.Insert(2))
	assert.False(t, board2.MoveRight())
}
