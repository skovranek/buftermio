package buftermio

import (
	"bufio"
	"os"
	"strings"
)

func NewBuffer(prompt ...string) Buffer {
	newScanner := bufio.NewScanner(os.Stdin)
	newScanner.Split(bufio.ScanBytes)
	newBuffer := Buffer{
		scanner: newScanner,
		buf:     make([]byte, 0),
		len:     0,
		cursor:  0,
		cache:   make([][]byte, 0),
		index:   0,
		prompt:  strings.Join(prompt, ""),
	}
	return newBuffer
}
