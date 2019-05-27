package main

import (
	"github.com/Cguilliman/terminal-file-browser/display"
	ui "github.com/gizak/termui"
)

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
				return
			case "<Up>":
				display.ListUp()
			case "<Down>":
				display.ListDown()
			default:
				eventID := e.ID
				charDescript(eventID, searchChan)
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
			case "<C-f>":
				searchChan := make(chan string)
				go display.InputProcess(searchChan)
				WriteHandle(display, searchChan)
				uiEvents = ui.PollEvents() // KOSTIL`
			case "<Up>":
				display.ListUp()
			case "<Down>":
				display.ListDown()
			case "<Enter>":
				display.SelectDir()
			}
		}
	}
}
