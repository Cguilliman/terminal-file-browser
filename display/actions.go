package display

import (
	ui "github.com/gizak/termui"
)

func defaultHandlers(display *Display, value string) bool {
	switch value {
	case "<Up>":
		display.Content.ListUp()
	case "<Down>":
		display.Content.ListDown()
	case "<PageDown>":
		display.Content.PageDown()
	case "<PageUp>":
		display.Content.PageUp()
	case "<Enter>":
		display.Content.SelectDir(false)
	default:
		return false
	}

	return true
}

func charDescript(char string, searchChan chan string) {
	if char == "<Space>" {
		char = " "
	}
	searchChan <- char
}

func WriteHandle(display *Display, charChan chan string) {
	uiEvents := ui.PollEvents()

	for e := range uiEvents {
		if e.Type == ui.KeyboardEvent {
			switch e.ID {
			case "<C-f>", "<C-q>":
				display.Content.Reset()
				display.SearchInput.Reset()
				close(charChan)
				return
			case "<Enter>":
				if display.Submit(charChan) {
					return
				}
			default:
				if !defaultHandlers(display, e.ID) {
					eventID := e.ID
					charDescript(eventID, charChan)
				}
			}
		}
	}
}

func ActionsHandle(display *Display) {
	uiEvents := ui.PollEvents()

	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "<C-q>":
				return
			case "<C-r>":
				charChan := display.Run()
				WriteHandle(display, charChan)
				uiEvents = ui.PollEvents() // KOSTIL`
			case "<C-f>": // Searching
				charChan := display.Search()
				WriteHandle(display, charChan)
				uiEvents = ui.PollEvents() // KOSTIL`
			case "<C-t>": // Create file
				charChan := display.Touch()
				WriteHandle(display, charChan)
				uiEvents = ui.PollEvents() // KOSTIL`
			case "<C-n>": // Make directory
				charChan := display.MkDir()
				WriteHandle(display, charChan)
				uiEvents = ui.PollEvents() // KOSTIL`
			case "<Backspace>":
				display.Content.SelectDir(true)
			default:
				defaultHandlers(display, e.ID)
			}
		}
	}
}
