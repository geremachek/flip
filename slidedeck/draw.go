package slidedeck

import "github.com/gdamore/tcell"

// add text to the screen

func addString(s tcell.Screen, style tcell.Style, x int, y int, text string) {
	for i, ch := range text {
		s.SetContent(i+x, y, ch, []rune{}, style)
	}
}