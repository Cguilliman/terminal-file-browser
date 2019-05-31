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

func WriteHandle(display *Display, charChan chan string, uiEvents <-chan ui.Event) {
	// uiEvents := ui.PollEvents()

	for e := range uiEvents {
		if e.Type == ui.KeyboardEvent {
			switch e.ID {
			case "<C-f>", "<C-q>":
				display.Content.Reset(true)
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
			case "<C-q>":  // quit
				return
			case "<Backspace>": // select directory
				display.Content.SelectDir(true)

			case "<Space>": // pick out file
				display.Content.PickOut()
			case "<C-a>": // pick out all files in directory
				display.Content.PickOutAll()
			case "<C-c>": // copy file or directory
				display.Content.Copy()
			case "<C-x>": // cut file or directory
				display.Content.Cut()
			case "<C-v>": // paste file or directory
				display.Content.Paste()
			case "<Delete>": // delete file or directory
				display.Content.Delete()

			case "<C-r>": // run command
				charChan := display.Run()
				WriteHandle(display, charChan, uiEvents)

			case "<C-f>": // search in directory
				charChan := display.Search()
				WriteHandle(display, charChan, uiEvents)

			case "<C-t>": // create file
				charChan := display.Touch()
				WriteHandle(display, charChan, uiEvents)

			case "<C-n>": // make directory
				charChan := display.MkDir()
				WriteHandle(display, charChan, uiEvents)

			case "<C-z>": // zip file or dir
				charChan := display.Zip()
				WriteHandle(display, charChan, uiEvents)

			case "<C-u>": // unzip file
				charChan := display.Unzip()
				WriteHandle(display, charChan, uiEvents)

			default:
				defaultHandlers(display, e.ID)
				continue
			}
		}
	}
}
