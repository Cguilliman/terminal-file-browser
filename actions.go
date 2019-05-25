package main

import (
	ui "github.com/gizak/termui"
)

func charDescript(char string, searchChan chan string) {
	if char == "<Space>" {
		char = " "
	}
	searchChan <- char
}

func WriteHandle(display *Display, searchChan chan string) {
	uiEvents := ui.PollEvents()

	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "<C-s>", "<C-q>":
				return
			}

			switch e.Type {
			case ui.KeyboardEvent:
				eventID := e.ID
				charDescript(eventID, searchChan)
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
			case "<C-s>":
				searchChan := make(chan string)
				go display.InputProcess(searchChan)
				WriteHandle(display, searchChan)
				uiEvents = ui.PollEvents()  // KOSTIL`
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
