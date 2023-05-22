package main

import (
	"fmt"
)

// REPL
func main() {
	buffer := NewBuffer()

	fmt.Println("enter 'q' to exit")

	loop := true
	for ; loop; {
		fmt.Print("$ ") // remove
		input := buffer.GetInput()

		fmt.Printf("Input: %s\n", input)

		if input == "q" {
			loop = false
			fmt.Println("QUIT")
		}
	}
}