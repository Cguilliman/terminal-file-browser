package main

import (
	"github.com/Cguilliman/terminal-file-browser/display"
	ui "github.com/gizak/termui"
)

func defaultHandlers(display *display.Display, value string) bool {
	switch value {
	case "<Up>":
		display.ListUp()
	case "<Down>":
		display.ListDown()
	case "<PageDown>":
		display.PageDown()
	case "<PageUp>":
		display.PageUp()
	case "<Enter>":
		display.SelectDir()
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

func WriteHandle(display *display.Display, searchChan chan string) {
	uiEvents := ui.PollEvents()

	for e := range uiEvents {
		if e.Type == ui.KeyboardEvent {
			switch e.ID {
			case "<C-f>", "<C-q>":
				display.ResetList()
				// TODO reset input widget
				return
			case "<Enter>":
				display.SelectDir()
				// TODO reset input widget
				return
			default:
				if !defaultHandlers(display, e.ID) {
					eventID := e.ID
					charDescript(eventID, searchChan)
				}
			}
		}
	}
}

func ActionsHandle(display *display.Display) {
	uiEvents := ui.PollEvents()

	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "<C-q>":
				return
			case "<C-f>":  // Searching
				charChan := display.SearchInputProcess()
				WriteHandle(display, charChan)
				uiEvents = ui.PollEvents() // KOSTIL`
			default:
				defaultHandlers(display, e.ID)
			}
		}
	}
}
