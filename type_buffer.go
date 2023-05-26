package buftermio

import (
	"bufio"
	"os"
	"strings"
)

// key byte slices, used as constants (not immutable)
var DELETE []byte = []byte{uint8(127)}
var UPARROW []byte = []byte{uint8(27), uint8(91), uint8(65)}
var DOWNARROW []byte = []byte{uint8(27), uint8(91), uint8(66)}
var RIGHTARROW []byte = []byte{uint8(27), uint8(91), uint8(67)}
var LEFTARROW []byte = []byte{uint8(27), uint8(91), uint8(68)}
var TAB []byte = []byte{uint8(9)}
var FOURSPACES ]byte = []byte{uint8(32), uint8(32), uint8(32), uint8(32)}
var ESCAPE ]byte = []byte{uint8(27)}
var OPENBRACKET []byte = []byte{uint8(91)}
var CARRIAGERETURN []byte = []byte{uint8(10)}

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
	NewBuffer(...string) Buffer
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

func NewBuffer(prompt ...string) Buffer {
	newScanner := bufio.NewScanner(os.Stdin)
	newScanner.Split(bufio.ScanBytes)
	newBuffer := Buffer{
		scanner: newScanner,
		buf:    make([]byte, 0),
		len:    0,
		cursor: 0,
		cache:  make([][]byte, 0),
		index:  0,
		prompt: strings.Join(prompt, ""),
	}
	return newBuffer
}