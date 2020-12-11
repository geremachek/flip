package slidedeck

import "github.com/gdamore/tcell"

func (sd *SlideDeck) Handle(key *tcell.EventKey) bool {
	if key.Key() == tcell.KeyRune {
		switch key.Rune() {
			case 'q': return false
			case 'b': sd.ToggleBar()
			case 'h': sd.MoveSlide(-1)
			case 'l': sd.MoveSlide(1)
		}

		sd.Redraw()
	} else {
		switch key.Key() {
			case tcell.KeyLeft: sd.MoveSlide(-1)
			case tcell.KeyRight: sd.MoveSlide(1)
		}

		sd.Redraw()
	}

	return true
}
