package tetris

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBoard(t *testing.T) {
	board, _ := NewBoard(5, 10, 2)
	assert.Equal(t, 5, board.width)
	assert.Equal(t, 10, board.height)
	assert.False(t, board.BlockExists(0, 0))
	assert.False(t, board.BlockExists(4, 0))
	assert.False(t, board.BlockExists(0, 9))
	assert.False(t, board.BlockExists(4, 9))

	// ボードの範囲外を指定してもFalseを返す
	assert.False(t, board.BlockExists(-1, 0))
	assert.False(t, board.BlockExists(0, -1))
	assert.False(t, board.BlockExists(4, 10))
	assert.False(t, board.BlockExists(5, 10))

	assert.Equal(t, 5, len(board.blocks))
	assert.Equal(t, 10, len(board.blocks[0]))
}

func TestBlockController(t *testing.T) {
	board, err := NewBoard(5, 2, 2)
	assert.Nil(t, err)
	controller := NewBlockController(board)
	assert.True(t, controller.Insert())
	assert.Equal(t, 2, controller.x)
	assert.Equal(t, 1, controller.y)
	assert.True(t, board.BlockExists(2, 1))
}

func TestMoveBlock(t *testing.T) {
	board, err := NewBoard(5, 2, 2)
	assert.Nil(t, err)
	assert.True(t, board.Insert())
	assert.True(t, board.BlockExists(2, 1))
	assert.Equal(t, 2, board.moveX)
	assert.Equal(t, 1, board.moveY)

	assert.True(t, board.MoveLeft())
	assert.False(t, board.BlockExists(2, 1))
	assert.True(t, board.BlockExists(1, 1))

	assert.True(t, board.MoveBottom())
	assert.False(t, board.BlockExists(1, 1))
	assert.True(t, board.BlockExists(1, 0))

	assert.True(t, board.MoveRight())
	assert.False(t, board.BlockExists(1, 0))
	assert.True(t, board.BlockExists(2, 0))
	assert.True(t, board.isMove)

	assert.True(t, board.MoveBottom())
	assert.True(t, board.BlockExists(2, 0))
	assert.False(t, board.isMove)
	assert.False(t, board.IsGameOver())

	//ブロック二つ目
	board.Insert()
	assert.True(t, board.isMove)
	assert.False(t, board.IsGameOver())
	assert.True(t, board.MoveBottom())
	assert.False(t, board.isMove)
	assert.True(t, board.IsGameOver())

}

func TestInsertBlockMiss(t *testing.T) {
	//範囲外
	board, err := NewBoard(5, 10, 6)
	assert.Nil(t, board)
	assert.Error(t, err)
	board, err = NewBoard(5, 10, -1)
	assert.Nil(t, board)
	assert.Error(t, err)

	//投入場所に既にに存在する
	board, _ = NewBoard(5, 10, 0)
	assert.True(t, board.Insert())
	assert.False(t, board.Insert())
	//制御中に投入する
	assert.False(t, board.Insert())
}
func TestMoveBlockMiss(t *testing.T) {
	board, _ := NewBoard(3, 10, 0)
	//範囲外
	assert.True(t, board.Insert())
	assert.False(t, board.MoveLeft())

	board2, _ := NewBoard(3, 2, 2)
	//範囲外
	assert.True(t, board2.Insert())
	assert.False(t, board2.MoveRight())

	//接地後に動かそうとしたら失敗
	board3, _ := NewBoard(3, 2, 1)
	board3.Insert()
	board3.MoveBottom()
	board3.MoveBottom()
	assert.False(t, board3.MoveLeft())
	assert.False(t, board3.MoveRight())
	assert.False(t, board3.MoveBottom())

	// 左右移動先にブロックがあったら動かせない
	board4, _ := NewBoard(3, 2, 1)
	board4.Insert()
	board4.MoveLeft()
	board4.MoveBottom()
	board4.MoveBottom()
	board4.Insert()
	board4.MoveLeft()
	board4.MoveBottom()
	board4.MoveBottom()

	board4.Insert()
	board4.MoveRight()
	board4.MoveBottom()
	board4.MoveBottom()
	board4.Insert()
	board4.MoveRight()
	board4.MoveBottom()
	board4.MoveBottom()

	board4.Insert()
	assert.False(t, board4.MoveLeft())
	assert.False(t, board4.MoveRight())
}
