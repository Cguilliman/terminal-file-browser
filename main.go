package main

import (
	"flag"
	"fmt"
	"github.com/Cguilliman/terminal-file-browser/display"
	mg "github.com/Cguilliman/terminal-file-browser/manager"
	ui "github.com/gizak/termui"
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
	defer ui.Close()

	// infinite processing loop
	display := display.InitDisplay(manager)
	ActionsHandle(display)
}
