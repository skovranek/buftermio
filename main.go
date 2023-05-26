package main

import (
	"fmt"
	"os"
)

// REPL
func main() {
	buffer := NewBuffer("Enter: ")

	fmt.Println("Enter 'q' to exit.")

	for {
		input, err := buffer.GetInput()

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Input: %s\n", input)
			fmt.Printf("Bytes: %v\n\n", []byte(input))
			if input == "q" {
				fmt.Print("QUIT\n\n")
				os.Exit(0)
			}
		}
	}
}