package buftermio

import (
	"bufio"
	"os"
	"strings"
)

func NewBuffer(defaultPrompt ...string) Buffer {
	newScanner := bufio.NewScanner(os.Stdin)
	newScanner.Split(bufio.ScanBytes)
	newBuffer := Buffer{
		scanner: newScanner,
		line:    cachedBuffer{},
		cache:   make([]cachedBuffer, 0),
		index:   0,
		prompt:  strings.Join(defaultPrompt, ""),
	}
	return newBuffer
}
