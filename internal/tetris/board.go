package tetris

type Board struct {
	width, height int
	blocks [][]bool
}

func NewBoard(width, height int) *Board {
	b := new(Board)
	b.width = width
	b.height = height
	b.blocks = make([][]bool, width)
	for column := range b.blocks {
		b.blocks[column] = make([]bool, height)
		for row := range b.blocks {
			b.blocks[column][row] = false
		}
	}
	return b
}

func (b *Board) blockExists(width, height int) bool {
	return b.blocks[width][height]
}

func (b *Board) put(width, height int) bool {
	if b.width > width && b.height > height {
		for i :=0; b.height > height + i; i++ {
			if !b.blockExists(width, height + i) {
				b.blocks[width][height + i] = true
				return true
			}
		}
	}
	return false
}
