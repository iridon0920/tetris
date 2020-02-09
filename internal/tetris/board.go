package tetris

import "errors"

type Board struct {
	width, height int
	blocks        [][]bool
	isMove        bool
	moveX, moveY  int
	insertX       int
	HeadY         int
}

func NewBoard(width, height, insertX int) (*Board, error) {
	if width <= insertX || 0 > insertX {
		return nil, errors.New("ブロック挿入位置が不正です。")
	}

	b := new(Board)
	b.width = width
	b.height = height
	b.HeadY = b.height - 1
	b.insertX = insertX
	b.blocks = make([][]bool, width)
	for column := range b.blocks {
		b.blocks[column] = make([]bool, height)
		for row := range b.blocks[column] {
			b.blocks[column][row] = false
		}
	}
	return b, nil
}

func (b *Board) BlockExists(x, y int) bool {
	return b.blocks[x][y]
}

func (b *Board) IsGameOver() bool {
	return b.BlockExists(b.insertX, b.HeadY) && !b.isMove
}

func (b *Board) Insert() bool {
	if !b.BlockExists(b.insertX, b.HeadY) && !b.isMove {
		b.blocks[b.insertX][b.HeadY] = true
		b.moveX = b.insertX
		b.moveY = b.HeadY
		b.isMove = true
		return true
	}
	return false
}

func (b *Board) MoveLeft() bool {
	if b.moveX-1 >= 0 && b.isMove && !b.BlockExists(b.moveX-1, b.moveY) {
		b.blocks[b.moveX][b.moveY] = false
		b.moveX--
		b.blocks[b.moveX][b.moveY] = true
		return true
	}
	return false
}

func (b *Board) MoveRight() bool {
	if b.moveX+1 < b.width && b.isMove && !b.BlockExists(b.moveX+1, b.moveY) {
		b.blocks[b.moveX][b.moveY] = false
		b.moveX++
		b.blocks[b.moveX][b.moveY] = true
		return true
	}
	return false
}

func (b *Board) MoveBottom() bool {
	if b.isMove {
		if b.moveY == 0 || b.BlockExists(b.moveX, b.moveY-1) {
			b.isMove = false
		} else {
			b.blocks[b.moveX][b.moveY] = false
			b.moveY--
			b.blocks[b.moveX][b.moveY] = true
		}
		return true
	}
	return false
}
