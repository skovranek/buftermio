package main

import (
	"fmt"
	exec "golang.org/x/sys/execabs"
)

// REPL
func main() {
	defer func(){exec.Command("stty", "-f", "/dev/tty", "sane").Run()}()

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