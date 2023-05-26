package buftermio

import exec "golang.org/x/sys/execabs"

func PrepTerm() {
	// turn off buffer
	exec.Command("stty", "-f", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display Stdin
	exec.Command("stty", "-f", "/dev/tty", "-echo").Run()
}

func DeferSane() {
	// set terminal to normal
	exec.Command("stty", "-f", "/dev/tty", "sane").Run()
}