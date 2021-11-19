package main

import (
	"os"
	"fmt"
	"flag"
	sd "github.com/geremachek/flip/slidedeck"
)

func main() {
	var (
		stdin = flag.Bool("s", false, "Read slides from stdin")
		slides [][]string
	)

	flag.Parse()

	args := flag.Args()

	if *stdin { // our slides are coming from stdin
		slides = processSlide(os.Stdin)
	} else if len(args) > 0 { // if there are files supplied as arguments...
		// read them!
		
		var err error

		if slides, err = getDeck(args); err != nil {
			printError("couldn't read file(s)")
		}
	} else { // otherwise print an error
		printError("files must be supplied as arguments")
	}

	// start the interface blah blah blah...

	if deck, err := sd.NewSlideDeck(slides, true); err == nil {
		if err := deck.Start(); err != nil {
			printError("couldn't start interface")
		}
	} else {
		printError("couldn't get screen")
	}

}

func printError(msg string) {
	fmt.Fprintf(os.Stderr, "flip: %s\n", msg)
	os.Exit(1)
}
