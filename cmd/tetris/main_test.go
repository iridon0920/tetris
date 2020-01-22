package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	exp := "ここからテトリスを構築していく。"
	got := Main()
	if exp != got {
		t.Errorf("exp: %v got:%v",exp,got)
	}
}
