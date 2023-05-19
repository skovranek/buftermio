package main

import (
	"fmt"
)

// REPL
func main() {
	defer DeferMe()

	ch := make(chan string)

	loop := true
	
	for ; loop; {
		go GetInput(ch)
		input := <- ch
		fmt.Printf("Input: %s\n", input)
		if input == "q" {
			loop = false
			fmt.Println("QUIT")
		}
	}
}