package slidedeck

import (
	"fmt"
	"github.com/gdamore/tcell"
	"github.com/geremachek/flip/draw"
)

// CheckResize
// Resize if it the terminal is out of sync with the program

func (sd *SlideDeck) Resize(w, h int) {
	sd.maxWidth = w

	realHeight := h

	if sd.barVisible { // we need to make room for the bar
		realHeight -= 2
	}

	sd.maxHeight = realHeight

	sd.Redraw()
}

// DrawFile
// Draw a file to the screen, wrapping lines

func (sd SlideDeck) DrawFile() {
	var fitLines []string

	for _, line := range sd.deck[sd.card] {
		if len(line) > sd.maxWidth { // the line is too long to be displayed
			fitLines = append(fitLines, line[:sd.maxWidth])
			fitLines = append(fitLines, line[sd.maxWidth:])
		} else {
			fitLines = append(fitLines, line)
		}
	}

	for y, line := range fitLines {
		if y < sd.maxHeight {
			draw.Addstr(sd.screen, tcell.StyleDefault, 0, y, line)
		} else {
			break
		}
	}
}

// MoveSlide
// Move a slide, either left or right

func (sd *SlideDeck) MoveSlide(n int) {
	var (
		edge int
		start int
		len int = len(sd.deck) - 1
	)

	if n == 1 {
		edge = len
		start = 0
	} else if n == -1 {
		edge = 0
		start = len
	}

	if sd.card == edge {
		sd.card = start
	} else {
		sd.card += n
	}
}

// DrawBar
// Draw a status bar below the displayed file, if the bar is set to visible

func (sd *SlideDeck) DrawBar() {
	if sd.barVisible {
		style := tcell.StyleDefault.Underline(true)
		info := fmt.Sprintf("%d/%d", sd.card + 1, len(sd.deck))

		w, h := sd.screen.Size()

		draw.Addstr(sd.screen, style, w-len(info), h-1, info)
	}
}

// ToggleBar
// Taggle the status bar

func (sd *SlideDeck) ToggleBar() {
	sd.barVisible = !sd.barVisible

	if !sd.barVisible { // we want to hide the bar
		sd.maxHeight += 2
	} else { // we want to show the bar
		sd.maxHeight -= 2
	}
}

// Redraw
// Clear, and redraw the screen

func (sd *SlideDeck) Redraw() {
	sd.screen.Clear()
	
	sd.DrawFile()
	sd.DrawBar()

	sd.screen.Show()
}
