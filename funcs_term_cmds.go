package buftermio

import (
	"log"

	exec "golang.org/x/sys/execabs"
)

func prepTerm() {
	// turn off buffer
	if err := exec.Command("stty", "-f", "/dev/tty", "cbreak", "min", "1").Run(); err != nil {
		log.Fatal(err)
	}

	// do not display Stdin
	if err := exec.Command("stty", "-f", "/dev/tty", "-echo").Run(); err != nil {
		log.Fatal(err)
	}
}

func deferSane() {
	// set terminal to normal
	if err := exec.Command("stty", "-f", "/dev/tty", "sane").Run(); err != nil {
		log.Fatal(err)
	}
}
