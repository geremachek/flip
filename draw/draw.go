package draw

import "github.com/gdamore/tcell"

// Addstr
// add text to the screen

func AddRunes(s tcell.Screen, style tcell.Style, x int, y int, runes []rune) {
	for i := x; i < len(runes)+x; i++ {
		s.SetContent(i, y, runes[i-x], []rune(""), style)
	}
}
