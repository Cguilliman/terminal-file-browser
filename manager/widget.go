package manager

import (
    ui "github.com/gizak/termui"
    "github.com/gizak/termui/widgets"
)

type Widget struct {
    renderChan chan []string
    widget *widgets.List 
}

func (self *Widget) Render(isRender bool) {
    ui.Render(self.widget)
}

func (self *Widget) ScrollUp() {
    self.widget.ScrollUp()
    self.Render(true)
}

func (self *Widget) ScrollDown() {
    self.widget.ScrollDown()
    self.Render(true)
}

func (self *Widget) PageUp() {
    self.widget.ScrollPageUp()
    self.Render(true)
}

func (self *Widget) PageDown() {
    self.widget.ScrollPageDown()
    self.Render(true)
}

func (self *Widget) SelectDir(rows []string) {
    self.widget.ScrollTop()
    self.widget.Rows = rows
    self.Render(true)
}

func (self *Widget) renderLoop() {
    for list := range self.renderChan {
        self.widget.Rows = list
        self.widget.ScrollTop()
        self.Render(true)
    }
}

func initWidget() (*Widget, chan []string) {
    renderChan := make(chan []string)
    
    obj := Widget{renderChan, widgets.NewList()}
    obj.widget.SetRect(0, 3, 80, 20)
    obj.widget.TextStyle = ui.NewStyle(ui.ColorYellow)

    go obj.renderLoop()
    return &obj, renderChan
}
