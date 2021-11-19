package draw

import "github.com/gdamore/tcell"

// Addstr
// add text to the screen

func AddString(s tcell.Screen, style tcell.Style, x int, y int, text string) {
	curs := x

	for _, ch := range text {
		s.SetContent(curs, y, ch, []rune{}, style)
		curs++
	}
}
