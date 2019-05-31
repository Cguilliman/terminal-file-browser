package display

import (
	"github.com/Cguilliman/terminal-file-browser/inputs"
	"github.com/Cguilliman/terminal-file-browser/manager"
)

const (
	SEARCH string = "search"
	RUN    string = "run"
)

type Display struct {
	Content      *manager.ContentList
	SearchInput  *inputs.Input
	RunInput     *inputs.Input
	currentFocus string
}

func (self *Display) Submit(charChan chan string) bool {
	if self.currentFocus == SEARCH {
		self.Content.SelectDir(false)
		close(charChan)
		return true
	} else if self.currentFocus == RUN {
		close(charChan)
		return true
	}
	return false
}

func (self *Display) Search() chan string {
	self.currentFocus = SEARCH
	charChan := make(chan string)
	searchChan := make(chan string)

	go self.SearchInput.InputProcess(charChan, searchChan, false)
	go self.Content.SearchProcess(searchChan)

	return charChan
}

func (self *Display) Run() chan string {
	return Command(
		self,
		func(runChan chan string, content *manager.ContentList) {
			manager.Run(runChan, content.Manager)
		},
	)
}

func (self *Display) MkDir() chan string {
	return Command(self, CreateDir)
}

func (self *Display) Touch() chan string {
	return Command(self, CreateFile)
}

func (self *Display) Zip() chan string {
	return Command(self, manager.Zipping)
}

func (self *Display) Unzip() chan string {
	return Command(self, manager.Unzipping)
}

func InitDisplay() *Display {
	content := manager.Init("")
	searchInput := inputs.Init("", [4]int{0, 0, 80, 3})

	display := &Display{
		Content:      content,
		SearchInput:  searchInput,
		currentFocus: "",
	}
	return display
}
