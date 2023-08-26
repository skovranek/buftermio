package buftermio

import "fmt"

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

		if b.contains(del) {
			b.backspace()
		} else if b.contains(upArrow) {
			b.upIndex()
		} else if b.contains(downArrow) {
			b.downIndex()
		} else if b.contains(rightArrow) {
			b.cursorRight()
		} else if b.contains(leftArrow) {
			b.cursorLeft()
		} else if b.contains(tab) {
			b.fourSpaces()
        } else if b.contains(ctrlA) {
            b.cursorSOL()
        } else if b.contains(ctrlE) {
            b.cursorEOL()
        //} else if b.contains(ctrlW) {
        //    b.deleteWord()
		// if byte is escape key, skip it (to handle arrow bytes)
		} else if b.contains(esc) {
			prevEsc = true
		// if byte is part of arrow's escape sequence, skip it
		} else if b.contains(openBracket) && prevEsc {
			prevEsc = false
		} else {
			fmt.Print(string(next))
			if b.cursor < b.len {
				fmt.Print(string(b.buf[b.cursor:]))
				left(b.len - b.cursor)
			}
			prevEsc = false
		}
		if b.contains(carriageReturn) {
			return b.enter(), nil
		}
	}
	return "", nil
}
