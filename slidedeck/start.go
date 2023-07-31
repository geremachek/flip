package slidedeck

import "github.com/gdamore/tcell"

func (sd *slideDeck) Start() error {
	if err := sd.screen.Init(); err == nil {
		var (
			input tcell.Event
			running bool = true
		)

		for running {
			input = sd.screen.PollEvent()

			switch ev := input.(type) {
				case *tcell.EventKey:    running = sd.handle(ev) 
				case *tcell.EventResize: sd.resize()
			}
		}

		sd.screen.Fini()
	} else {
		return err
	}

	return nil
}
