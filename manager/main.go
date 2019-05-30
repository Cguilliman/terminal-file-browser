package manager

type ContentListInterface interface {
	ListUp()
	PageUp()
	ListDown()
	PageDown()
	SelectDir()
}

type ContentList struct {
	RenderChan chan UpdateData
	Widget     *Widget
	Manager    *Manager
}

func (self *ContentList) ListUp() {
	self.RenderChan <- UpdateData{
		self.Manager.PrevFile(),
		"UP", "",
	}
}

func (self *ContentList) ListDown() {
	self.RenderChan <- UpdateData{
		self.Manager.NextFile(),
		"DOWN", "",
	}
}

func (self *ContentList) PageUp() {
	self.RenderChan <- UpdateData{
		self.Manager.SetFirstFile(),
		"PAGEUP", "",
	}
}

func (self *ContentList) PageDown() {
	self.RenderChan <- UpdateData{
		self.Manager.SetLastFile(),
		"PAGEDOWN", "",
	}
}

func (self *ContentList) Reset() {
	self.Manager.SetFiles()
	self.RenderChan <- UpdateData{
		self.Manager.RenderList(nil), 
		"", "",
	}
}

func (self *ContentList) SelectDir(isParent bool) {
	list, err := self.Manager.EnterDir(isParent)

	if err != nil {
		return
	}
	self.RenderChan <- UpdateData{
		list, "GOTOP", 
		self.Manager.Path,
	}
}

func (self *ContentList) SearchProcess(searchChan chan string) {
	self.Manager.Search(searchChan, self.RenderChan)
	self.Widget.GoTop()
}

func Init(path string) *ContentList {
	var content ContentList                           // init contentList
	content.Widget, content.RenderChan = initWidget() // init widget and re-render channel
	content.Manager = initManager(path)               // init manager worker
	content.RenderChan <- UpdateData{                 // push current files rows in channel
		content.Manager.RenderList(nil), 
		"", content.Manager.Path,
	} 
	return &content
}
