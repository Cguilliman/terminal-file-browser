package main

import (
	tm "github.com/nsf/termbox-go"
	ui "github.com/gizak/termui"
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
