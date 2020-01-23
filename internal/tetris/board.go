package tetris

type board struct{
	width int
	height int
	blocks [][]string
}

func NewBoard(width int, height int) *board {
	board := new(board)
	board.width = width
	board.height = height

	blocks := make([][]string, height)
	for i := 0; i < height; i++ {
		line := make([]string, width)
		for j:=0; j < width; j++ {
			line[j] = "."
		}
		blocks[i] = line
	}
	board.blocks = blocks
	return board
}

func (board *board) OutputRow(row int) string {
	result := ""
	for _, v := range board.blocks[row] {
		result += v
	}
	return result
}
