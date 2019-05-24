package main

import (
	"fmt"
	mg "github.com/Cguilliman/terminal-file-browser/manager"
	ui "github.com/gizak/termui"
	"github.com/gizak/termui/widgets"
)

var (
	_ = fmt.Println
)

type Display struct {
	// Header
	List *widgets.List
	// Footer
	Manager *mg.Manager
}

func (display *Display) ListUp() {
	display.List.ScrollUp()    // more `cursor`
	display.Manager.PrevFile() // change current file position
	display.InitList()         // re-render files list
	ui.Render(display.List)
}

func (display *Display) ListDown() {
	display.List.ScrollDown()  // move `cursor`
	display.Manager.NaxtFile() // change current file position
	display.InitList()         // re-render files list
	ui.Render(display.List)
}

func (display *Display) SelectDir() {
	// NOTE: mb I need to remove it to another place
	// TODO: add invalidation `if`
	err := display.Manager.EnterDir() // change directory and get files list
	if err != nil {
		return
	}
	display.InitList()         // re-initialize list of files in display
	display.List.ScrollTop()
	ui.Render(display.List)
}

func (display *Display) InitList() {
	// mb some more list customization
	display.List.Title = display.Manager.Path
	display.List.Rows = display.Manager.RenderList()
}

func InitDisplay(manager *mg.Manager) *Display {
	list := widgets.NewList()
	defer ui.Render(list)

	list.WrapText = true
	list.SetRect(0, 0, 80, 20)
	list.TextStyle = ui.NewStyle(ui.ColorYellow)

	display := &Display{list, manager}
	display.InitList()
	return display
}
