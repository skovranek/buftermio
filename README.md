# buftermio
buffer for terminal input/output --> __*buf*__ + __*term*__ + __*io*__

[![](https://godoc.org/github.com/skovranek/buftermio?status.svg)](https://pkg.go.dev/github.com/skovranek/buftermio#section-readme)
## What
Buftermio is an importable Go package. It provides an cached buffer for input from the CLI. Perfect for a REPL!

_Basically, an improved interface for a bufio.Scanner reading os.Stdin._
## Why
When you use a CLI shell, you may take for granted that you have a cached history of your commands. Then, when you go to use your own CLI program, suddenly you've lost that cache and have to type each command anew instead of just pressing up to scroll through previous commands. Here is buftermio to the rescue. Buftermio lets you scroll through your previous inputs with the up and down arrow keys!
## How
Inspired by the bytes package [Buffer](https://pkg.go.dev/bytes#Buffer) struct. Buftermio uses a [bufio.Scanner](https://pkg.go.dev/bufio#Scanner) to read from [os.Stdin](https://pkg.go.dev/os#pkg-variables) for input. It temporarily prevents the CLI shell from printing directly to Stdout while buftermio reads from Stdin. Buftermio intercepts each key to interpret them before printing them to Stdout.
## Download/Install
Run
```
go get github.com/skovranek/buftermio
```
## Configure
I made this to be as simple and straight forward as possible. No configuration except the optional prompt like a shell would have. See the NewBuffer function: [GitHub](https://github.com/skovranek/buftermio/blob/main/type_buffer.go) or [GoDoc](https://pkg.go.dev/github.com/skovranek/buftermio#NewBuffer)
## Implement
1) Instantiate the buffer _**once**_ with the NewBuffer function from buftermio. It takes optional variatic string arguments which are joined to form the prompt which prints out before getting input.
```go
buffer := buftermio.NewBuffer("Hello ", username, ":")
// or
buffer := buftermio.NewBuffer("$ ")
// or
buffer := buftermio.NewBuffer()
```
2) Call the GetInput method on the buffer each time you want to get input from the CLI. GetInput returns a string.
```go
input, err := buffer.GetInput()
```
## Example
Here is an example of a REPL that uses buftermio to echo the input and print the keycodes. Useful for checking keycodes.
```go
package main

import (
    "fmt"

    "github.com/skovranek/buftermio"
)

func main() {
    cont := true
    buffer := buftermio.NewBuffer("input: ")
	
    for {
        input, err := buffer.GetInput()
        if err != nil {
            fmt.Println(err)
            continue
        }

        if input == "q" {
            return
        }

        fmt.Printf("output: %s\n", input)
        fmt.Printf("bytes: %v\n\n", []byte(input))
    }
}
```
## UI Controls
Using buftermio should feel like using the interface of a shell like bash or zsh.
- Up/down arrows to scroll though the cached history of inputs.
- Left/right arrows to move the cursor left/right. Buftermio will insert characters instead of overwriting them.
- Return to enter input.
- Tab will output four spaces. I found the tab functionality was inconsistent in the CLI, so I simplified it.
- Ctrl+A will move the cursor to the start of the line.
- Ctrl+E will move the cursor to the end of the line.
- Ctrl+W will delete the previous word.
> **WARNING**
> Pressing keys or key combinations that produce keycode sequences, such as <Alt+A>, may not work as expected. This is a known issue and may cause errors because buftermio is not a keylogger counting keystrokes. It is only counting the character bytes from Stdin. A-Z, 0-9, and most keys will behave normally. But, if on the off chance there are unpredicted effects when you use certain keys, you now know why. I'll gladly add checks for certain sequences for anyone that requires it.
## Dependencies
Buftermio uses a subrepository from the Go Project called [execabs](https://pkg.go.dev/golang.org/x/sys/execabs) instead of the standard library's "os/exec" package because of a [path-security issue](https://go.dev/blog/path-security) with "os/exec". No other dependencies are included.
## Testing
Manually tested with a zsh shell in the macOS terminal. Mocking stdin would only test the mock. Please let me know if buftermio does not work as expected in your environment.
## Contact
Questions, issues or suggestions: mattjskov at gmail.com
## Contribute
Feel free to chip in. Let's work together to customize buftermio for your project. Submit pull requests to the 'main' branch.
