package manager

import (
	ui "github.com/gizak/termui"
	"github.com/gizak/termui/widgets"
)

var (
	COMMANDS map[string]func(*Widget) = map[string]func(*Widget){
		"UP":       func(widget *Widget) { widget.Up() },
		"DOWN":     func(widget *Widget) { widget.Down() },
		"PAGEUP":   func(widget *Widget) { widget.PageUp() },
		"PAGEDOWN": func(widget *Widget) { widget.PageDown() },
		"GOTOP":    func(widget *Widget) { widget.GoTop() },
	}
)

type UpdateData struct {
	list []string
	command string
}

type Widget struct {
	renderChan chan UpdateData
	widget     *widgets.List
}

func (self *Widget) Render(isRender bool) {
	ui.Render(self.widget)
}

func (self *Widget) Up() {
	self.widget.ScrollUp()
}

func (self *Widget) Down() {
	self.widget.ScrollDown()
}

func (self *Widget) PageUp() {
	self.widget.ScrollPageUp()
}

func (self *Widget) PageDown() {
	self.widget.ScrollPageDown()
}

func (self *Widget) GoTop() {
	self.widget.ScrollTop()
}

func (self *Widget) SelectDir(rows []string) {
	self.GoTop()
	self.widget.Rows = rows
	self.Render(true)
}

func (self *Widget) renderLoop() {
	for item := range self.renderChan {
		if item.command != "" {
			self.runCommand(item.command)
		}
		self.widget.Rows = item.list
		self.Render(true)
	}
}

func (self *Widget) runCommand(command string) {
	// TODO add command validation
	COMMANDS[command](self)
}

func initWidget() (*Widget, chan UpdateData) {
	renderChan := make(chan UpdateData)

	obj := Widget{renderChan, widgets.NewList()}
	obj.widget.SetRect(0, 3, 80, 20)
	obj.widget.TextStyle = ui.NewStyle(ui.ColorYellow)

	go obj.renderLoop()
	return &obj, renderChan
}
