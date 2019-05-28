package main

import (
	"flag"
	"fmt"
	dp "github.com/Cguilliman/terminal-file-browser/display"
	ui "github.com/gizak/termui"
)

func printHelp() {
	fmt.Println("super help")
}

func main() {
	// obtain and processing flags
	var (
		helpFlag = flag.Bool("h", false, "display help dialog")
		// pathFlag = flag.String("path", "", "path to open directory")
	)
	if *helpFlag {
		printHelp()
		return
	}
	
	// initialize GUI
	if err := ui.Init(); err != nil {
		panic(err)
	}
	defer ui.Close()

	display := dp.InitDisplay()
	// infinite processing loop
	dp.ActionsHandle(display)
}
