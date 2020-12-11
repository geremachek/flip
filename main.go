package main

import (
	"os"
	"fmt"
	"github.com/gdamore/tcell"
	sd "github.com/geremachek/flip/slidedeck"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		printError("files must be supplied as arguments")
	} else {
		if d, err := GetDeck(args); err == nil {
			if s, err := tcell.NewScreen(); err == nil {
				if err = s.Init(); err == nil {
					deck := sd.NewSlideDeck(s, d, true)

					var (
						input tcell.Event
						running bool = true
					)

					deck.DrawFile()
					deck.DrawBar()

					s.Show()

					for running {
						input = s.PollEvent()

						switch ev := input.(type) {
							case *tcell.EventKey: running = deck.Handle(ev) 
							case *tcell.EventResize:
								w, h := ev.Size()
								deck.Resize(w, h)
						}
					}

					s.Fini()
				} else {
					printError("couldn't initialize screen")
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
