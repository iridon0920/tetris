package tetris

type BlockController struct {
	x, y  int
	board *Board
}

func NewBlockController(board *Board) *BlockController {
	c := new(BlockController)
	c.x = board.insertX
	c.y = board.HeadY
	c.board = board
	return c
}
