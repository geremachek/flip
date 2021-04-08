package main

import (
	"os"
	"fmt"
	sd "github.com/geremachek/flip/slidedeck"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		printError("files must be supplied as arguments")
	} else {
		if d, err := getDeck(args); err == nil {
			if deck, err := sd.NewSlideDeck(d, true); err == nil {
				if err := deck.Start(); err != nil {
					printError("couldn't start interface")
				}
			} else {
				printError("couldn't get screen")
			}
		} else {
			printError("couldn't read file(s)")
		}
	}
}

func printError(msg string) {
	fmt.Fprintln(os.Stderr, "flip: " + msg)
}
