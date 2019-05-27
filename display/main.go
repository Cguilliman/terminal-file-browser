package display

import (
	"fmt"
	mg "github.com/Cguilliman/terminal-file-browser/manager"
	ui "github.com/gizak/termui"
	"github.com/gizak/termui/widgets"
)

var (
	_ = fmt.Println
)

type DisplayInterface interface {
	ListUp()
	PageUp()
	ListDown()
	PageDown()
	SelectDir()
	ResetList()
	SearchInputProcess()
}

type Display struct {
	List    *widgets.List
	Input   *Input
	Manager *mg.Manager
}

func (self *Display) ListUp() {
	self.List.ScrollUp()    // more `cursor`
	self.Manager.PrevFile() // change current file position
	self.initList(true)     // re-render files list
}

func (self *Display) ListDown() {
	self.List.ScrollDown()  // move `cursor`
	self.Manager.NextFile() // change current file position
	self.initList(true)     // re-render files list
}

func (self *Display) PageUp() {
	self.List.ScrollPageUp()    // more `cursor`
	self.Manager.FirstFile() // change current file position
	self.initList(true)     // re-render files list
}

func (self *Display) PageDown() {
	self.List.ScrollPageDown()  // move `cursor`
	self.Manager.LastFile() // change current file position
	self.initList(true)     // re-render files list
}

func (self *Display) SelectDir() {
	err := self.Manager.EnterDir() // change current directory and get files list
	if err != nil {
		return
	}
	self.initList(false) // re-initialize list of files in display
	self.List.ScrollTop()
	ui.Render(self.List)
}

func (self *Display) ResetList() {
	self.Manager.SetFiles()
	self.initList(true)
}

func (self *Display) initList(isRerender bool) {
	if isRerender {
		defer ui.Render(self.List)
	}
	self.List.Title = self.Manager.Path
	self.List.Rows = self.Manager.RenderList(nil)
}

func (self *Display) rerenderLoop(filesChan chan []string) {
	for list := range filesChan {
		self.List.Rows = list
		ui.Render(self.List)
	}
}

func (self *Display) SearchInputProcess() chan string {
	charChan := make(chan string)
	searchChan := make(chan string)
	renderChan := make(chan []string)

	go self.Input.InputProcess(charChan, searchChan)
	go self.Manager.Search(searchChan, renderChan)
	go self.rerenderLoop(renderChan)

	return charChan
}

func InitDisplay(manager *mg.Manager) *Display {
	list := widgets.NewList()
	input := InitInput(false)
	defer ui.Render(list, input.Widget)

	list.WrapText = true
	list.SetRect(0, 3, 80, 20)
	list.TextStyle = ui.NewStyle(ui.ColorYellow)

	display := &Display{list, input, manager}
	display.initList(false)
	return display
}
