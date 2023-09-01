package buftermio

import (
	"fmt"
	"strings"
)

func (b *Buffer) backspace() {
	if b.line.cursor > 1 {
		b.removeSlice(2)
		printLeft(b.line.cursor + 1)
		fmt.Print(string(b.line.buf), " ")
		printLeft(b.line.len - b.line.cursor + 1)
	} else {
		b.removeSlice(1)
		bell()
	}
}

func (b *Buffer) upIndex() {
	b.removeSlice(3)
	if b.index > 0 {
		b.addIndex(b.index)
		b.clear()
		b.index--
		if b.index > 0 && b.isIndexEmpty(b.index) {
			b.index--
		}
		b.newLine(b.cache[b.index])
	} else {
		bell()
	}
}

func (b *Buffer) downIndex() {
	b.removeSlice(3)
	if b.index < len(b.cache)-1 {
		b.addIndex(b.index)
		b.clear()
		b.index++
		b.newLine(b.cache[b.index])
	} else {
		bell()
	}
}

func (b *Buffer) cursorRight() {
	b.removeSlice(3)
	if b.line.cursor < b.line.len {
		fmt.Print(string(right))
		b.line.cursor++
	} else {
		bell()
	}
}

func (b *Buffer) cursorLeft() {
	b.removeSlice(3)
	if b.line.cursor > 0 {
		printLeft(1)
		b.line.cursor--
	} else {
		bell()
	}
}

func (b *Buffer) cursorSOL() {
	b.removeSlice(1)
	if b.line.cursor > 0 {
		printLeft(b.line.cursor)
		b.line.cursor = 0
	} else {
		bell()
	}
}

func (b *Buffer) cursorEOL() {
	b.removeSlice(1)
	if b.line.cursor < b.line.len {
		fmt.Print(strings.Repeat(string(right), b.line.len-b.line.cursor))
		b.line.cursor = b.line.len
	} else {
		bell()
	}
}

func (b *Buffer) deleteWord() {
	b.removeSlice(1)
	if b.line.cursor == 0 {
		bell()
		return
	}
	isWord := false
	for {
		if b.line.cursor > 0 {
			if isWord && b.line.buf[b.line.cursor-1] == space {
				return
			}
			if !isWord && b.line.buf[b.line.cursor-1] != space {
				isWord = true
			}
			b.removeSlice(1)
			printLeft(b.line.cursor + 1)
			fmt.Print(string(b.line.buf), " ")
			printLeft(b.line.len - b.line.cursor + 1)
		} else {
			return
		}
	}
}

func (b *Buffer) deleteSOL() {
	b.removeSlice(1)
	if b.line.cursor > 0 {
		printLeft(b.line.cursor)
		fmt.Print(strings.Repeat(" ", b.line.len))
		printLeft(b.line.len)
		fmt.Print(string(b.line.buf[b.line.cursor:]))
		printLeft(b.line.len - b.line.cursor)
		b.line.len = b.line.len - b.line.cursor
		b.line.buf = b.line.buf[b.line.cursor:]
		b.line.cursor = 0
	} else {
		bell()
	}
}

func (b *Buffer) deleteEOL() {
	b.removeSlice(1)
	if b.line.cursor < b.line.len {
		fmt.Print(strings.Repeat(" ", b.line.len-b.line.cursor))
		printLeft(b.line.len - b.line.cursor)
		b.line.len = b.line.cursor
		b.line.buf = b.line.buf[:b.line.cursor]
	} else {
		bell()
	}
}

func (b *Buffer) fourSpaces() {
	b.removeSlice(1)
	b.insert(fourSpaces)
	printLeft(b.line.cursor - 4)
	fmt.Print(string(b.line.buf))
	printLeft(b.line.len - b.line.cursor)
}

func (b *Buffer) enter() string {
	b.removeSlice(1)
	if b.index < len(b.cache) {
		newCache := make([]cachedBuffer, 0)
		newCache = append(newCache, b.cache[:b.index]...)
		newCache = append(newCache, b.cache[b.index+1:]...)
		b.cache = newCache
	}
	if b.line.len > 0 {
		b.addIndex(len(b.cache))
	}
	b.index = len(b.cache)
	copy := make([]byte, 0)
	copy = append(copy, b.line.buf...)
	b.line = cachedBuffer{}
	b.line.len = 0
	return string(copy)
}
