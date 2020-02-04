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
