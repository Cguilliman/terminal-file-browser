package manager

import (
	tmp "github.com/Cguilliman/terminal-file-browser/manager/temporary"
)

type ContentList struct {
	RenderChan chan UpdateData
	Widget     *Widget
	Manager    *Manager
	tempFiles  *tmp.TemporaryFiles // contain files witch will be remove/copy
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

func (self *ContentList) PickOut() {
	self.Manager.PickOutFile()
	self.RenderChan <- UpdateData{
		self.Manager.RenderList(nil),
		"", "",
	}
}

func (self *ContentList) PickOutAll() {
	self.Manager.PickOutAllFiles()
	self.RenderChan <- UpdateData{
		self.Manager.RenderList(nil),
		"", "",
	}
}

func (self *ContentList) GetSelectedFiles() []string {
	var files []string
	for _, item := range self.Manager.Highlighting {
		files = append(
			files,
			self.Manager.GetDirPath(item),
		)
	}
	if len(files) == 0 && self.Manager.CurrentFileNumber > 1 {
		// set current file/directory in focus
		files = []string{self.Manager.GetDirPath(-1)}
	}
	return files
}

func (self *ContentList) addTemporary(action string) {
	self.tempFiles = &tmp.TemporaryFiles{
		TemporaryFiles: self.GetSelectedFiles(),
		Action:         action,
		BaseDirectory:  self.Manager.Path,
	}
}

func (self *ContentList) Cut() {
	self.addTemporary(tmp.CUT)
}

func (self *ContentList) Copy() {
	self.addTemporary(tmp.COPY)
}

func (self *ContentList) Delete() {
	self.addTemporary(tmp.DEL)
	self.tempFiles.Delete()

	self.Manager.SetFiles()
	self.RenderChan <- UpdateData{
		self.Manager.SetFirstFile(),
		"GOTOP", "",
	}
}

func (self *ContentList) Paste() {
	if self.tempFiles != nil {
		self.tempFiles.Paste(self.Manager.Path)
	}
	self.Reset(true)
}

func (self *ContentList) Reset(isHighlight bool) {
	if isHighlight {
		self.Manager.DelHighlighting()
	}
	self.Manager.SetFiles()
	self.RenderChan <- UpdateData{
		self.Manager.RenderList(nil),
		"", "",
	}
}

func (self *ContentList) SelectDir(isParent bool) {
	self.Manager.DelHighlighting()
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
	self.Manager.DelHighlighting()
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
	content.tempFiles = nil
	return &content
}
