package buftermio

import (
	"bytes"
	"fmt"
	"strings"
)

func (b *Buffer) insert(data []byte) {
	copy := make([]byte, 0)
	copy = append(copy, b.buf[:b.cursor]...)
	copy = append(copy, data...)
	copy = append(copy, b.buf[b.cursor:]...)
	b.buf = copy
	b.len += len(data)
	b.cursor += len(data)
}

func (b *Buffer) contains(data []byte) bool {
	return bytes.Index(b.buf, data) >= 0
}

func (b *Buffer) newLine(data ...byte) {
	b.buf = make([]byte, 0)
	b.buf = append(b.buf, data...)
	b.len = len(data)
	b.cursor = len(data)
}

func (b *Buffer) removeSlice(n int) {
	if b.cursor - n >= 0 {
		copy := make([]byte, 0)
		copy = append(copy, b.buf[:b.cursor - n]...)
		copy = append(copy, b.buf[b.cursor:]...)
		b.buf = copy
		b.len -= n
		b.cursor -= n
	}
}

func (b *Buffer) clear() {
	left(b.cursor) // fmt.Print(strings.Repeat(string(leftArrow), b.cursor))
	fmt.Print(strings.Repeat(" ", b.len))
	left(b.len)    // fmt.Print(strings.Repeat(string(leftArrow), b.len))
}

func (b *Buffer) isIndexEmpty(n int) bool {
	return len(bytes.TrimSpace(b.cache[n])) == 0
}

func (b *Buffer) addIndex(n int) {
	// if buffer is empty and previous cache entry is empty, skip add
	if b.len == 0 && b.isIndexEmpty(n-1) {
		return
	// else if n is not last index, replace at index n
	} else if n < len(b.cache) {
		b.cache[n] = b.buf
	/// or if last index is empty
	} else if len(b.cache) > 0 && b.isIndexEmpty(len(b.cache)-1) {
		b.cache[len(b.cache)-1] = b.buf
	} else {
		b.cache = append(b.cache, b.buf)
	}
}
