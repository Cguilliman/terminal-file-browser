package manager

type ContentListInterface interface {
    ListUp()
    PageUp()
    ListDown()
    PageDown()
    SelectDir()
}

type ContentList struct {
    RenderChan chan []string
    Widget *Widget
    Manager *Manager
}

func (self *ContentList) ListUp() {
    self.Manager.PrevFile()
    self.Widget.ScrollUp()
}

func (self *ContentList) ListDown() {
    self.Manager.NextFile()
    self.Widget.ScrollDown()
}

func (self *ContentList) PageUp() {
    self.Manager.SetFirstFile()
    self.Widget.PageUp()
}

func (self *ContentList) PageDown() {
    self.Manager.SetLastFile()
    self.Widget.PageDown()
}

func (self *ContentList) Reset() {
    self.Manager.SetFiles()
    self.RenderChan <- self.Manager.RenderList(nil)
}

func (self *ContentList) SelectDir() {
    list, err := self.Manager.EnterDir()
    if err != nil {
        return
    }
    self.RenderChan <- list
}

func (self *ContentList) SearchProcess(searchChan chan string) {
    go self.Manager.Search(searchChan, self.RenderChan)
}

func Init(path string) *ContentList {
    var content ContentList  // init contentList
    content.Widget, content.RenderChan = initWidget()  // init widget and re-render channel 
    content.Manager = initManager(path)  // init manager worker
    content.RenderChan <- content.Manager.RenderList(nil)  // push current files rows in channel
    return &content
}
