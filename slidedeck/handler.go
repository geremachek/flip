package slidedeck

import "github.com/gdamore/tcell"

func (sd *slideDeck) handle(key *tcell.EventKey) bool {
	switch key.Key() {
		case tcell.KeyLeft:   sd.moveSlide(-1)
		case tcell.KeyRight:  sd.moveSlide(1)
		case tcell.KeyRune:
			switch key.Rune() {
				case 'q': return false
				case 'b': sd.toggleBar()
				case 'h': sd.moveSlide(-1)
				case 'l': sd.moveSlide(1)
			}
	}

	sd.redraw()
	
	return true
}
