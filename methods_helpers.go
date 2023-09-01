package buftermio

import (
	"bytes"
	"fmt"
	"strings"
)

func (b *Buffer) insert(data []byte) {
	copy := make([]byte, 0)
	copy = append(copy, b.line.buf[:b.line.cursor]...)
	copy = append(copy, data...)
	copy = append(copy, b.line.buf[b.line.cursor:]...)
	b.line.buf = copy
	b.line.len += len(data)
	b.line.cursor += len(data)
}

func (b *Buffer) newLine(cachedBuf cachedBuffer) {
	//b.line = cachedBuffer{}
	//b.line.buf = append(b.line.buf, data...)
	//b.line.cursor = len(data)
	b.line = cachedBuf
	b.line.len = len(cachedBuf.buf)
	fmt.Print(string(b.line.buf))
	printLeft(b.line.len - b.line.cursor)
}

func (b *Buffer) removeSlice(n int) {
	if b.line.cursor-n >= 0 {
		copy := make([]byte, 0)
		copy = append(copy, b.line.buf[:b.line.cursor-n]...)
		copy = append(copy, b.line.buf[b.line.cursor:]...)
		b.line.buf = copy
		b.line.len -= n
		b.line.cursor -= n
	}
}

func (b *Buffer) clear() {
	printLeft(b.line.cursor) // fmt.Print(strings.Repeat(string(printLeft), b.line.cursor))
	fmt.Print(strings.Repeat(" ", b.line.len))
	printLeft(b.line.len) // fmt.Print(strings.Repeat(string(printLeft), b.len))
}

func (b *Buffer) isIndexEmpty(n int) bool {
	return len(bytes.TrimSpace(b.cache[n].buf)) == 0
}

func (b *Buffer) addIndex(n int) {
	// if line.buffer is empty and previous cache entry is empty, skip add
	if b.line.len == 0 && b.isIndexEmpty(n-1) {
		return
		// else if n is not last index, replace at index n
	} else if n < len(b.cache) {
		b.cache[n] = b.line
		/// or if last index is empty
	} else if len(b.cache) > 0 && b.isIndexEmpty(len(b.cache)-1) {
		b.cache[len(b.cache)-1] = b.line
	} else {
		b.cache = append(b.cache, b.line)
	}
}
