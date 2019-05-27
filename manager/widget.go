package manager

import (
    ui "github.com/gizak/termui"
    "github.com/gizak/termui/widgets"
)

type Widget struct {
    widget *widgets.List 
}

func (self *Widget) Render(isRender bool) {

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
    self.widget.PageUp()
    self.Render(true)
}

func (self *Widget) PageDown() {
    self.widget.PageDown()
    self.Render(true)
}

func (self *Widget) SelectDir(rows []string) {
    self.widget.ScrollTop()
    self.widget.Rows = rows
    self.Render(true)
}

func (self *Widget) renderLoop(renderChan chan []string) {
    for list := range renderChan {
        self.List.Rows = list
        self.List.ScrollTop()
        self.Render(true)
    }
}
