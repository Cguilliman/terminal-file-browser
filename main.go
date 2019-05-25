package main

import (
	"flag"
	"fmt"
	mg "github.com/Cguilliman/terminal-file-browser/manager"
	ui "github.com/gizak/termui"
	tm "github.com/nsf/termbox-go"
)

func printHelp() {
	fmt.Println("super help")
}

func main() {
	// obtain and processing flags
	var (
		helpFlag = flag.Bool("h", false, "display help dialog")
		pathFlag = flag.String("path", "", "path to open directory")
	)
	if *helpFlag {
		printHelp()
		return 
	}

	// initialize manager
	manager := mg.InitManager(*pathFlag)

	// initialize GUI
	if err := ui.Init(); err != nil {
		panic(err)
	}
	// tm.SetInputMode(tm.InputEsc)
	defer Shutdown()

	// infinite processing loop
	display := InitDisplay(manager)
	ActionsHandle(display)
	// for {
	// 	if exit {
	// 		return
	// 	}
	// }
}

func Shutdown() {
	if tm.IsInit {
		ui.Close()
	}
}
