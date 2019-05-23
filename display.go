package main

import (
	"fmt"
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
	Manager *Manager
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
	display.Manager.SelectDir() // change directory and get files list
	display.InitList()          // re-initialize list of files in display
	ui.Render(display.List)
}

func (display *Display) InitList() {
	// mb some more list customization
	display.List.Rows = display.Manager.RenderList()
}

func InitDisplay(manager *Manager) *Display {
	list := widgets.NewList()
	defer ui.Render(list)

	list.Title = manager.Path
	list.WrapText = true
	list.SetRect(0, 0, 80, 20)
	list.TextStyle = ui.NewStyle(ui.ColorYellow)

	display := &Display{list, manager}
	display.InitList()
	return display
}
