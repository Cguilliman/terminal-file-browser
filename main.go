package main

import (
	ui "github.com/gizak/termui"
	tm "github.com/nsf/termbox-go"
)

func main() {
	manager := InitManager()
	_ = manager

	if err := ui.Init(); err != nil {
		panic(err)
	}
	tm.SetInputMode(tm.InputEsc)
	defer Shutdown()

	display := InitDisplay(manager)

	for {
		exit := ActionsHandle(display)
		if exit {
			return
		}
	}
}

func Shutdown() {
	if tm.IsInit {
		ui.Close()
	}
}
