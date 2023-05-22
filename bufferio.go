package main

import (
	exec "golang.org/x/sys/execabs"
	"bufio"
	"os"
	"bytes"
	"fmt"
)

/*
note: cursor in shell dictates where chars print
TODO:
[x] remove right/left arrow bytes
[x] add cursor property
[ ] ++/-- cursor position as needed
[ ] frontspace: backspace in front of cursor
[ ] overwrite byte slice at cursor
[ ] overwrite byte slice at index (& cursor)
[ ] add custom prompt $ 
*/

var carriageReturn []byte = []byte{uint8(10)}
var deleteByte     []byte = []byte{uint8(127)}
var upArrow        []byte = []byte{uint8(27), uint8(91), uint8(65)}
var downArrow      []byte = []byte{uint8(27), uint8(91), uint8(66)}
var rightArrow     []byte = []byte{uint8(27), uint8(91), uint8(67)}
var leftArrow      []byte = []byte{uint8(27), uint8(91), uint8(68)}

type buffer struct {
	reader  *bufio.Scanner
	buf     []byte
	start   int
	end     int
	cursor  int
	history [][]byte
	index   int
}

type bufferI interface {
	// trying a new indenting style for readability
	 NewBuffer() buffer
       readLen() int
      indexLen() int
          read() []byte
         print()
      contains([]byte) bool
         write(...byte)
       newLine(...byte)
    removeLast(int)
    	   add()
      GetInput() string
}

func NewBuffer() buffer {
	newReader := bufio.NewScanner(os.Stdin)
	newReader.Split(bufio.ScanBytes)
	newBuffer := buffer{
		reader:  newReader,
		buf:     make([]byte, 0),
		start:   0,
		end:     0,
		cursor:  0,
		history: make([][]byte, 0),
		index:   0,
	}
	return newBuffer
}

func (b *buffer) readLen() int {
	return b.end - b.start
}

func (b *buffer) indexLen(n int) int {
	return len(bytes.TrimSpace(b.history[n]))
}

func (b *buffer) read() []byte {
	// use copy instead??? TODO
	return b.buf[b.start: b.end]
}

func (b *buffer) print() {
	fmt.Print(string(b.read()))
}

func (b *buffer) contains(data []byte) bool {
	if i := bytes.Index(b.read(), data); i >= 0 {
		return true
	}
	return false
}

func (b *buffer) write(data ...byte) {
	b.buf = append(b.buf, data...)
	b.end = len(b.buf)
}

func (b *buffer) newLine(data ...byte) {
	next := len(b.buf)
	b.write(data...)
	b.start = next
}

func (b *buffer) removeLast(n int) {
	if b.readLen() >= n {
		b.end -= n
		b.newLine(b.read()...)
	}
}

func (b *buffer) add() {
	if len(b.history) > 0 && b.indexLen(len(b.history)-1) == 0 {
		b.history[len(b.history)-1] = b.read()
	} else {
		b.history = append(b.history, b.read())
	}
}

func (b *buffer) GetInput() string {
	PrepTerm()
	defer DeferSane()

	for b.reader.Scan() {
		newByte := b.reader.Bytes()
		b.write(newByte...)
		fmt.Print(string(newByte))

		if b.contains(carriageReturn) {
			b.removeLast(1)
			if b.readLen() > 0 {
				b.add()
				b.index = len(b.history)
			}
			copy := make([]byte, 0)
			copy = append(copy, b.read()...)
			b.newLine()
			return string(copy)

		} else if b.contains(deleteByte) {
			if b.end - b.start > 1 {
				b.removeLast(2)
				backSpace(1)
			} else {
				b.removeLast(1)
				log(7)
			}

		} else if b.contains(upArrow) {
			b.removeLast(3)
			log(27, 91, 66)
			if b.index == len(b.history) && b.index > 0 {
				b.add()
			}
			if b.index > 0 {
				backSpace(b.readLen())
				b.index--
				b.newLine(b.history[b.index]...)
				b.print()
			} else {
				log(7) // bell
			}

		} else if b.contains(downArrow) {
			b.removeLast(3)
			if b.index < len(b.history)-1 {
				backSpace(b.readLen())
				b.index++
				b.newLine(b.history[b.index]...)
				b.print()
			} else {
				log(7) // bell
			}

		} else if b.contains(rightArrow) {
			b.removeLast(3)

		} else if b.contains(leftArrow) {
			b.removeLast(3)
		}
	}
	return ""
}

// Helpers

func log(n ...int) {
	for _, val := range n {
		fmt.Print(string([]byte{uint8(val)}))
	}
}

func backSpace(n int) {
	for i := 0; i < n; i++ {
		log(27, 91, 68, 32, 27, 91, 68) // LEFT, space, LEFT
	}
}

// Terminal commands

func PrepTerm() {
	// turn off buffer
	exec.Command("stty", "-f", "/dev/tty", "cbreak", "min", "1").Run()
	// do not print
	exec.Command("stty", "-f", "/dev/tty", "-echo").Run()
}

func DeferSane() {
	exec.Command("stty", "-f", "/dev/tty", "sane").Run()
}

/*
up:     [27 91 65] string([]byte{uint8(27), uint8(91), uint8(65)})
down:   [27 91 66] string([]byte{uint8(27), uint8(91), uint8(66)})
right:  [27 91 67] string([]byte{uint8(27), uint8(91), uint8(67)})
left:   [27 91 68] string([]byte{uint8(27), uint8(91), uint8(68)})
escape: 27
[:      91
up:     65 (A)
down:   66 (B)
right:  67 (C)
left:   68 (D)
enter:  10
space:  32
tab:     9

		//splitter:
		// if byte is NL/CR/enter/'\n'/uint8(10)
			// remove NL byte
			// store buf to history
			// set b.index to length
			// send buf out
			// reset buf to zero
			// reset cursor
		if b.contains(carriageReturn) {
			b.removeLast(1)
			// if last saved byte slice in b.history is empty, replace it
			b.add()
			b.index = len(b.history)
			copy := make([]byte, 0)
			copy = append(copy, b.read()...)
			output <- string(copy)
			b.newLine()

		// else if byte is BS/DEL/uint8(127)
			// remove byte BS
			// decrement cursor
			// print (without add bytes) left arrow
			// print space (without add bytes)
			// print (without add bytes) left arrow
		} else if b.contains(deleteByte) {
			b.removeLast(2)
			if b.end > 1 {
				backSpace(1)
			}

		// else if bytes are UP/[]byte{uint8(27), uint8(91), uint8(65)}
			// remove UP bytes
			// remove 3 from cursor
			// counter printed sequence by printing DOWN
			// remove printed line
			// if b.index == length of b.history:
				// store b to b.history and decrement b.index
			// decrement b.index
			// set b to prev string (b.history[i])
			// set cursor to length of b/string
			// print b
		} else if b.contains(upArrow) {
			b.removeLast(3)
			log(27, 91, 66)
			if b.index == len(b.history) && b.index > 0 {
				b.add()
			}
			if b.index > 0 {
				backSpace(b.readLen())
				b.index--
				// copy from b.history
				b.newLine(b.history[b.index]...)
				b.print()
			} else {
				log(7) // bell
			}

		// else if bytes are DOWN/[]byte{uint8(27), uint8(91), uint8(66)}
			// remove DOWN bytes
			// if b.index < length of b.history
				// remove printed sequence
				// increment b.index
				// set b to next string (b.history[b.index])
				// set cursor to length of b/string
				// print b
		} else if b.contains(downArrow) {
			b.removeLast(3)
			if b.index < len(b.history)-1 {
				backSpace(b.readLen())
				b.index++
				b.newLine(b.history[b.index]...)
				b.print()
			} else {
				log(7) // bell
			}
		}
		// process LEFT/RIGHT arrows?
*/