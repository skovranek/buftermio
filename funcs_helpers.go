package buftermio

import (
	"fmt"
	"strings"
)

func bell() {
	fmt.Print(string(uint8(7)))
}

func left(n int) {
	fmt.Print(strings.Repeat(string(leftArrow), n))
}
