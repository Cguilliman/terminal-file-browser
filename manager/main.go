package manager

type ContentListInterface interface {
    ListUp()
    PageUp()
    ListDown()
    PageDown()
    SelectDir()
}

type ContentList struct {
    Widget *Widget
    Manager *Manager
}

func (self *ContentList) ListUp() {
    self.List.ScrollUp()
    self.Manager.PrevFile()
    self.Widget.ScrollUp()
    self.initList(true)
}

func (self *ContentList) ListDown() {
    self.List.ScrollDown()
    self.Manager.NextFile()
    self.initList(true)
}

func (self *ContentList) PageUp() {
    self.List.ScrollPageUp()
    self.Manager.FirstFile()
    self.initList(true)
}

func (self *ContentList) PageDown() {
    self.List.ScrollPageDown()
    self.Manager.LastFile()
    self.initList(true)
}

func (self *ContentList) SelectDir() {
    list, err := self.Manager.EnterDir()
    if err != nil {
        return
    }
    self.Widget.SelectDir(list)
}

func (self *ContentList) SearchProcess(searchChan chan string) {
    renderChan := make(chan []string)

    go self.Manager.Search(searchChan, renderChan)
    go self.rerenderLoop(renderChan)
}
