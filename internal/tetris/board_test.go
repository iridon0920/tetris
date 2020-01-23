package tetris

import (
	"testing"
)

func TestNewBoard(t *testing.T) {
	width := 5
	height := 10
	expWidth := 5
	expHeight := 10
	board := NewBoard(width, height)

	if (board.width != expWidth) {
		t.Errorf("exp: %d",expWidth)
	}
	if (board.height != expHeight){
		t.Errorf("exp: %d",expWidth)
	}
	if (board.OutputRow(3) != ".....") {
		t.Errorf("exp *****")
	}
}

// func TestSetBlock (t *testing.T) {
// 	data := map[string]struct {
// 		X int
// 		Y int
// 		Width int
// 		Hegiht int
// 		Exp [][]string
// 	}{
// 		"左隅へ置く" : {
// 			X :
// 		}
// 	}
// }
