package inputs

import (
    ui "github.com/gizak/termui"
    "github.com/gizak/termui/widgets"
)

type Widget struct {
    renderChan chan string
    isRun bool
    widget *widgets.Paragraph
}

func (self *Widget) Render() {
    ui.Render(self.widget)
}

func (self *Widget) renderLoop() {
    self.isRun = true
    for val := range self.renderChan {
        self.widget.Text = val
        self.Render()
    }
    self.isRun = false
}

func initWidget(size [4]int) (chan string, *Widget) {
    renderChan := make(chan string)
    obj := Widget{renderChan, false, widgets.NewParagraph()}
    
    if len(size) == 0 {
        obj.widget.SetRect(0, 0, 80, 3)
    } else {
        obj.widget.SetRect(
            size[0],
            size[1],
            size[2],
            size[3],
        )
    }

    go obj.renderLoop()
    return renderChan, &obj
}
