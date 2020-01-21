package cli

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

// Input 標準入力処理用構造体
type Input struct {
	scanner *bufio.Scanner
}

func (input Input) getLine() string {
	input.scanner.Scan()
	return input.scanner.Text()
}

func (input Input) getSlice(sep string) []string {
	input.scanner.Scan()
	return strings.Split(input.scanner.Text(), sep)
}

func (input Input) getIntLine() int {
	result, err := strconv.Atoi(input.getLine())
	if err != nil {
		fmt.Println(err)
	}
	return result
}

func (input Input) getIntSlice(sep string) []int {
	var intSlice []int
	strSlice := input.getSlice(sep)
	for _, s := range strSlice {
		i, err := strconv.Atoi(s)
		if err == nil {
			intSlice = append(intSlice, i)
		} else {
			fmt.Println(err)
		}
	}
	return intSlice
}
