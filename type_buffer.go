package buftermio

import (
	"bufio"
)

type Buffer struct {
	scanner *bufio.Scanner
	buf     []byte
	len     int
	cursor  int
	cache   [][]byte
	index   int
	prompt  string
}
