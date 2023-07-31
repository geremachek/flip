package main

import (
	"os"
	"fmt"
	sd "github.com/geremachek/flip/slidedeck"
)

func main() {
	var (
		slides [][]string
		args []string = os.Args[1:]
	)

	if len(args) == 0  { // our slides are coming from stdin
		slides = processSlide(os.Stdin)
	} else { // there are files supplied as arguments...
		// read them!
		
		var err error

		if slides, err = getDeck(args); err != nil {
			printError(err)
		}
	}

	// start the interface blah blah blah...

	if deck, err := sd.NewSlideDeck(slides, true); err == nil {
		if err := deck.Start(); err != nil {
			printError(err)
		}
	} else {
		printError(err)
	}

}

// print an error message

func printError(err error) {
	fmt.Fprintf(os.Stderr, "flip: %s\n", err)
}