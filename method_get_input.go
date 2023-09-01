package buftermio

import (
	"bytes"
	"fmt"
	"strings"
)

func (b *Buffer) GetInput(prompt ...string) (string, error) {
	defer deferSane()
	prepTerm()

	if len(prompt) > 0 {
		fmt.Print(strings.Join(prompt, ""))
	} else {
		fmt.Print(b.prompt)
	}

	prevEsc := false

	for b.scanner.Scan() {
		next := b.scanner.Bytes()

		if err := b.scanner.Err(); err != nil {
			return "", err
		}

		if b.line.cursor < b.line.len {
			b.insert(next)
		} else {
			b.line.buf = append(b.line.buf, next...)
			b.line.len++
			b.line.cursor++
		}

		if bytes.Contains(b.line.buf, del) {
			b.backspace()
		} else if bytes.Contains(b.line.buf, up) {
			b.upIndex()
		} else if bytes.Contains(b.line.buf, down) {
			b.downIndex()
		} else if bytes.Contains(b.line.buf, right) {
			b.cursorRight()
		} else if bytes.Contains(b.line.buf, left) {
			b.cursorLeft()
		} else if bytes.Contains(b.line.buf, tab) {
			b.fourSpaces()
		} else if bytes.Contains(b.line.buf, soh) {
			b.cursorSOL()
		} else if bytes.Contains(b.line.buf, enq) {
			b.cursorEOL()
		} else if bytes.Contains(b.line.buf, etb) {
			b.deleteWord()
		} else if bytes.Contains(b.line.buf, nak) {
			b.deleteSOL()
		} else if bytes.Contains(b.line.buf, vt) {
			b.deleteEOL()
			// if byte is escape key, skip it (to handle arrow bytes)
		} else if bytes.Contains(b.line.buf, esc) {
			prevEsc = true
			// if byte is part of arrow's escape sequence, skip it
		} else if bytes.Contains(b.line.buf, openBracket) && prevEsc {
			prevEsc = false
		} else {
			fmt.Print(string(next))
			if b.line.cursor < b.line.len {
				fmt.Print(string(b.line.buf[b.line.cursor:]))
				printLeft(b.line.len - b.line.cursor)
			}
			prevEsc = false
		}
		if bytes.Contains(b.line.buf, cr) {
			return b.enter(), nil
		}
	}
	return "", nil
}
