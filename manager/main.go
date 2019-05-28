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
	Widget     *Widget
	Manager    *Manager
}

func (self *ContentList) ListUp() {
	self.Widget.Up()
	self.RenderChan <- self.Manager.PrevFile()
}

func (self *ContentList) ListDown() {
	self.Widget.Down()
	self.RenderChan <- self.Manager.NextFile()
}

func (self *ContentList) PageUp() {
	self.Widget.PageUp()
	self.RenderChan <- self.Manager.SetFirstFile()
}

func (self *ContentList) PageDown() {
	self.Widget.PageDown()
	self.RenderChan <- self.Manager.SetLastFile()
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
	self.Widget.GoTop()
	self.RenderChan <- list
}

func (self *ContentList) SearchProcess(searchChan chan string) {
	go self.Manager.Search(searchChan, self.RenderChan)
}

func Init(path string) *ContentList {
	var content ContentList                               // init contentList
	content.Widget, content.RenderChan = initWidget()     // init widget and re-render channel
	content.Manager = initManager(path)                   // init manager worker
	content.RenderChan <- content.Manager.RenderList(nil) // push current files rows in channel
	return &content
}
