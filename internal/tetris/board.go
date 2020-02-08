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

func (b *Board) Put(x, y, width, height int) bool {
	if b.width > x && b.height > y {
		if y != 0 {
			for i := y; i >= 0; i-- {
				y = i
				if b.BlockExists(x, i-1) {
					break
				}
			}
		}
		for i := 0; b.height > y+i; i++ {
			if !b.BlockExists(x, y+i) {
				for h := 0; height > h; h++ {
					if b.height > y+i+h {
						b.blocks[x][y+i+h] = true
					} else {
						return false
					}
				}
				return true
			}
		}
	}
	return false
}

func (b *Board) Insert(x int) bool {
	if b.width > x && 0 <= x {
		if !b.BlockExists(x, 0) && !b.isMove {
			b.blocks[x][0] = true
			b.moveX = x
			b.moveY = 0
			b.isMove = true
			return true
		}
	}
	return false
}

func (b *Board) MoveLeft() bool {
	if b.moveX-1 >= 0 {
		b.blocks[b.moveX][b.moveY] = false
		b.moveX--
		b.blocks[b.moveX][b.moveY] = true
		return true
	}
	return false
}

func (b *Board) MoveRight() bool {
	if b.moveX+1 < b.width {
		b.blocks[b.moveX][b.moveY] = false
		b.moveX++
		b.blocks[b.moveX][b.moveY] = true
		return true
	}
	return false
}

func (b *Board) MoveBottom() bool {
	b.blocks[b.moveX][b.moveY] = false
	b.moveY++
	b.blocks[b.moveX][b.moveY] = true
	return true
}
