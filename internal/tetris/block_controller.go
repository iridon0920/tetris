package tetris

type BlockController struct {
	x, y    int
	isBlock bool
	board   *Board
}

func NewBlockController(board *Board) *BlockController {
	c := new(BlockController)
	c.board = board
	return c
}

func (c *BlockController) Insert() bool {
	if !c.board.BlockExists(c.board.insertX, c.board.HeadY) && !c.isBlock {
		c.x = c.board.insertX
		c.y = c.board.HeadY
		c.isBlock = true
		return true
	}
	return false
}

func (c *BlockController) MoveLeft() bool {
	if c.x-1 >= 0 && !c.board.BlockExists(c.x-1, c.y) {
		c.x--
		return true
	}
	return false
}

func (c *BlockController) MoveRight() bool {
	if c.x+1 < c.board.width && !c.board.BlockExists(c.x+1, c.y) {
		c.x++
		return true
	}
	return false
}

func (c *BlockController) MoveBottom() {
	if c.y == 0 || c.board.BlockExists(c.x, c.y-1) {
		c.board.blocks[c.x][c.y] = true
		c.isBlock = false
	} else {
		c.y--
	}
}
