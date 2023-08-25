# buftermio
Buffer for Terminal Input/Output --> buf + term + io
## [Go Package](https://pkg.go.dev/github.com/skovranek/buftermio#section-readme)

## 
This is a cached buffer for input from your CLI for your Go program. Perfect for a REPL!
_Basically, an improved interface for a bufio.Scanner reading os.Stdin._

## Why
When you use CLI shell, you may take for granted that you have a cached history of your inputs. Then, when you go to use your own CLI program, suddenly you've lost that cache and have to type each command anew instead of just pressing up to scroll through previous commands. Here is buftermio to the rescue. You can now use the ability to scroll through your previous inputs with the up and down arrow keys.

I made this to be as simple and straight forward as possible. No configuration except the optional prompt like a shell would have.

## How

Inspired by and based on the bytes package [Buffer](https://pkg.go.dev/bytes#Buffer) struct. Buftermio isn't a goroutine, but uses a for-loop with bufio.Scanner.Scan() on os.Stdin while getting input.

### Example:

1) Instantiate the buffer once with the NewBuffer function from buftermio. It takes optional variatic string arguments to form the prompt which prints out before getting input.
```go
buffer := buftermio.NewBuffer("Hello ", username, ":")
// or
buffer := buftermio.NewBuffer("$ ")
// or
buffer := buftermio.NewBuffer()
```
2) Call the GetInput method on the buffer each time you want to get input from the terminal.
```go
input, err := buffer.GetInput()
```
### Controls

--Uses up/down arrows to scroll though cache.

--Left/right arrows for cursor. Inserts characters instead of overwriting.

--Return to enter input.

--Tab is four spaces.

Note: Keys with escape sequences may not work. I'll gladly modify this if it becomes a problem for anyone.

## Who
Just myself so far. Contact for questions, issues or suggestions: mattjskov@gmail.com
