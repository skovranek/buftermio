package buftermio

import (
	"bytes"
	"fmt"
)

func (b *Buffer) GetInput() (string, error) {
	defer deferSane()
	prepTerm()

	fmt.Print(b.prompt)

	prevEsc := false

	for b.scanner.Scan() {
		next := b.scanner.Bytes()

		if err := b.scanner.Err(); err != nil {
			return "", err
		}

		if b.cursor < b.len {
			b.insert(next)
		} else {
			b.buf = append(b.buf, next...)
			b.len++
			b.cursor++
		}

		if bytes.Contains(b.buf, del) {
			b.backspace()
		} else if bytes.Contains(b.buf, up) {
			b.upIndex()
		} else if bytes.Contains(b.buf, down) {
			b.downIndex()
		} else if bytes.Contains(b.buf, right) {
			b.cursorRight()
		} else if bytes.Contains(b.buf, left) {
			b.cursorLeft()
		} else if bytes.Contains(b.buf, tab) {
			b.fourSpaces()
		} else if bytes.Contains(b.buf, soh) {
			b.cursorSOL()
		} else if bytes.Contains(b.buf, enq) {
			b.cursorEOL()
		} else if bytes.Contains(b.buf, etb) {
			b.deleteWord()
			// if byte is escape key, skip it (to handle arrow bytes)
		} else if bytes.Contains(b.buf, esc) {
			prevEsc = true
			// if byte is part of arrow's escape sequence, skip it
		} else if bytes.Contains(b.buf, openBracket) && prevEsc {
			prevEsc = false
		} else {
			fmt.Print(string(next))
			if b.cursor < b.len {
				fmt.Print(string(b.buf[b.cursor:]))
				printLeft(b.len - b.cursor)
			}
			prevEsc = false
		}
		if bytes.Contains(b.buf, carriageReturn) {
			return b.enter(), nil
		}
	}
	return "", nil
}
