package slidedeck

import "github.com/gdamore/tcell"

type SlideDeck struct {
	deck [][]string
	card int

	barVisible bool

	maxWidth int
	maxHeight int

	screen tcell.Screen
}

func NewSlideDeck(s tcell.Screen, d [][]string, bv bool) SlideDeck {
	w, h := s.Size()

	if bv {
		h -= 2; // make room for the bar
	}

	return SlideDeck { d, 0, bv, w, h, s }
}