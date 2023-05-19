package main

import (
	exec "golang.org/x/sys/execabs"
	"bufio"
	"os"
	"bytes"
	"fmt"
)

func GetInput(output chan string) {
	exec.Command("stty", "-f", "/dev/tty", "cbreak", "min", "1").Run()
	exec.Command("stty", "-f", "/dev/tty", "-echo").Run()
	//defer func(){exec.Command("stty", "-f", "/dev/tty", "sane").Run()}()

	reader := bufio.NewScanner(os.Stdin)
	reader.Split(bufio.ScanBytes)

	history := make([][]byte, 0)
	index := 0

	buffer := make([]byte, 0)
	cursor := 0
	
	logN(7)
	
	for reader.Scan() {
		b := reader.Bytes()
		buffer = append(buffer, b...)

		cursor++

		if i := bytes.IndexByte(buffer, uint8(10)); i >= 0 {
			buffer = removeLast(buffer, 1)
			// if last saved byte slice in history is empty, replace it
			if len(history) > 0 && len(history[len(history)-1]) == 0 {
				history[len(history)-1] = buffer
			} else {
				history = append(history, buffer)
			}
			index = len(history)
			copy := make([]byte, 0)
			copy = append(copy, buffer...)
			output <- string(copy)
			buffer = make([]byte, 0)
			cursor = 0

		} else if i := bytes.IndexByte(buffer, uint8(127)); i >= 0 {
			buffer = removeLast(buffer, 2)
			if cursor > 1 {
				cursor -= 2
				backSpace(1)
			}

		} else if i := bytes.Index(buffer, []byte{uint8(27), uint8(91), uint8(65)}); i >= 0 {
			buffer = removeLast(buffer, 3)
			cursor = len(buffer)
			logN(27, 91, 66)
			if index == len(history) && index > 0 {
				history = append(history, buffer)
			}
			if index > 0 {
				backSpace(len(buffer))
				index--
				buffer = make([]byte, 0)
				buffer = append(buffer, history[index]...)
				cursor = len(buffer)
				log(buffer)
			} else {
				logN(7) // bell
			}

		} else if i := bytes.Index(buffer, []byte{uint8(27), uint8(91), uint8(66)}); i >= 0 {
			cursor -= 3
			buffer = removeLast(buffer, 3)
			if index < len(history)-1 {
				backSpace(len(buffer))
				index++
				buffer = make([]byte, 0)
				buffer = append(buffer, history[index]...)
				cursor = len(buffer)
				log(buffer)
			} else {
				logN(7) // bell
			}
		}
		// for debugging to see when buffer goes to nil
		fmt.Println()
		log(buffer)
	}
}

func log(data []byte) {
	fmt.Print(string(data))
}

func logN(n ...int) {
	for _, val := range n {
		fmt.Print(string([]byte{uint8(val)}))
	}
}

func removeLast(data []byte, n int) []byte {
	if len(data) >= n {
		return data[0 : len(data)-n]
	}
	return data
}

func backSpace(n int) {
	for i := 0; i < n; i++ {
		logN(27, 91, 68, 32, 27, 91, 68) // LEFT, space, LEFT
	}
}

func DeferMe() {
	exec.Command("stty", "-f", "/dev/tty", "sane").Run()
}