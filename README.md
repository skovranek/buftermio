# buftermio
Buffer Terminal I/O --> Buf + ter + io

Basically an improved interface for a bufio.Scanner reading os.Stdin. 

This is a cached buffer for your terminal for doing I/O with your go CLI programs. Tested on mac terminal only.

Created this instead of using a bigger import like "term" because this package answers a very specific need.

Based on bytes package Buffer struct.

Operates like a shell with cached history of inputs. No commands like echo, etc. That is left open.

--Uses up/down to nav cache.

--Left/right arrows for cursor. Inserts characters instead of overwriting.

--Tab is four spaces.

--Keys with escape sequences may not work.

Not used as a goroutine, but uses a for loop with bufio.Scanner.Scan() on os.Stdin while getting input.

Uses go1.14.15 because my machine is a macbook running El Capitan.

example:
buffer := buftermio.NewBuffer("$ ")
input, err := buffer.GetInput()
