package main

import (
	tm "github.com/nsf/termbox-go"
)

func ActionsHandle(display *Display) bool {
	switch event := tm.PollEvent(); event.Key {
	case tm.KeyCtrlQ:
		return true
	case tm.KeyArrowUp:
		display.ListUp()
		return false
	case tm.KeyArrowDown:
		display.ListDown()
		return false
	case tm.KeyEnter:
		display.SelectDir()
		return false
		// case tm.KeyCtrlC:
		// 	CopyInBuffer(display.Manager.PathToCurrentFile())
		// 	return false
	}
	return false
}
