package tetris

type Board struct {
	width, height int
	blocks        [][]bool
	isMove        bool
	moveX, moveY  int
}

func NewBoard(width, height int) *Board {
	b := new(Board)
	b.width = width
	b.height = height
	b.blocks = make([][]bool, width)
	for column := range b.blocks {
		b.blocks[column] = make([]bool, height)
		for row := range b.blocks[column] {
			b.blocks[column][row] = false
		}
	}
	return b
}

func (b *Board) BlockExists(x, y int) bool {
	return b.blocks[x][y]
}

func (b *Board) Insert(x int) bool {
	head := b.height - 1
	if b.width > x && 0 <= x {
		if !b.BlockExists(x, head) && !b.isMove {
			b.blocks[x][head] = true
			b.moveX = x
			b.moveY = head
			b.isMove = true
			return true
		}
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
