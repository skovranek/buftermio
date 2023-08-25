package buftermio

import (
	"bufio"
)

type Buffer struct {
	scanner *bufio.Scanner
	buf    []byte
	len    int
	cursor int
	cache  [][]byte
	index  int
	prompt string
}

type buffer interface {
	GetInput() (string, error)
	backSpace()
	upIndex()
	downIndex()
	cursorRight()
	cursorLeft()
	fourSpaces()
	enter() string
	insert([]byte)
	contains([]byte) bool
	newLine(...byte)
	removeSlice(int)
	isIndexEmpty() bool
	addIndex(int)
}
