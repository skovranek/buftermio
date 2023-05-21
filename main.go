package main

import (
	"fmt"
)

// REPL
func main() {
	PrepTerm()

	defer DeferMe()

	buffer := NewBuffer()

	fmt.Println("enter 'q' to exit")

	loop := true
	for ; loop; {
		input := buffer.GetInput()

		fmt.Printf("Input: %s\n", input)

		if input == "q" {
			loop = false
			fmt.Println("QUIT")
		}
	}
}