package slidedeck

import "github.com/gdamore/tcell"

type slideDeck struct {
	deck [][]string
	card int

	barVisible bool

	maxWidth int
	maxHeight int

	screen tcell.Screen
}

func NewSlideDeck(d [][]string, bv bool) (*slideDeck, error) {
	if s, err := tcell.NewScreen(); err == nil {
		w, h := s.Size()

		if bv {
			h -= 2; // make room for the bar
		}

		return &slideDeck { d, 0, bv, w, h, s }, nil
	} else {
		return &slideDeck{}, err
	}
}
