package slidedeck

import (
	"fmt"
	"github.com/gdamore/tcell"
)

// Resize if it the terminal is out of sync with the program

func (sd *slideDeck) resize() {
	w, h := sd.screen.Size()

	sd.maxWidth = w
	realHeight := h

	if sd.barVisible { // we need to make room for the bar
		realHeight -= 2
	}

	sd.maxHeight = realHeight

	sd.redraw()
}

// Draw a file to the screen, wrapping lines

func (sd slideDeck) drawFile() {
	var x, y int

	// move the cursor down and send the cursor to the far left.

	goDown := func() bool {
		x = 0
		y++

		return y >= sd.maxHeight
	}

	for _, line := range sd.deck[sd.card] {
		// draw lines, wrapping if they're too long

		for _, ch := range line {
			sd.screen.SetContent(x, y, ch, []rune{}, tcell.StyleDefault)

			if x == sd.maxWidth-1 { // we are at the end of the line, kid.
				if goDown() { break } // go down! (break if not enough room)
			} else {
				x++ // otherwise just move the cursor to the right
			}
		}

		if goDown() { break }
	}
}

// Move a slide, either left or right

func (sd *slideDeck) moveSlide(n int) {
	var (
		len int = len(sd.deck) - 1
		edge int = len
		start int = 0
	)

	if n < 0 {
		edge = 0
		start = len
	}

	if sd.card == edge {
		sd.card = start
	} else {
		sd.card += n
	}
}

// Draw a status bar below the displayed file, if the bar is set to visible

func (sd *slideDeck) drawBar() {
	if sd.barVisible {
		style := tcell.StyleDefault.Underline(true)
		info := fmt.Sprintf("%d/%d", sd.card + 1, len(sd.deck))

		addString(sd.screen, style, 0, sd.maxHeight+1, info)
	}
}

// Taggle the status bar

func (sd *slideDeck) toggleBar() {
	sd.barVisible = !sd.barVisible

	if !sd.barVisible { // we want to hide the bar
		sd.maxHeight += 2
	} else { // we want to show the bar
		sd.maxHeight -= 2
	}
}

// Clear, and redraw the screen

func (sd *slideDeck) redraw() {
	sd.screen.Clear()
	
	sd.drawFile()
	sd.drawBar()

	sd.screen.Show()
}
