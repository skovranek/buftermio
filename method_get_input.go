package buftermio

import "fmt"

func (b *Buffer) GetInput() (string, error) {
	PrepTerm()
	defer DeferSane()

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

		if b.contains(DELETE) {
			b.backSpace()
		} else if b.contains(UPARROW) {
			b.upIndex()
		} else if b.contains(DOWNARROW) {
			b.downIndex()
		} else if b.contains(RIGHTARROW) {
			b.cursorRight()
		} else if b.contains(LEFTARROW) {
			b.cursorLeft()
		} else if b.contains(TAB) {
			b.fourSpaces()

		// if byte is escape key, skip it (to handle arrow bytes)
		} else if b.contains(ESCAPE) {
			prevEsc = true
		// if byte is part of arrow's escape sequence, skip it
		} else if b.contains(OPENBRACKET) && prevEsc {
			prevEsc = false
		} else {
			fmt.Print(string(next))
			if b.cursor < b.len {
				fmt.Print(string(b.buf[b.cursor:]))
				left(b.len - b.cursor)
			}
			prevEsc = false
		}
		if b.contains(CARRIAGERETURN) {
			return b.enter(), nil
		}
	}
	return "", nil
}