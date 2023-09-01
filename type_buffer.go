package buftermio

import (
	"bufio"
)

type Buffer struct {
	scanner *bufio.Scanner
	line    cachedBuffer
	cache   []cachedBuffer
	index   int
	prompt  string
}

type cachedBuffer struct {
	buf    []byte
	cursor int
	len    int
}
