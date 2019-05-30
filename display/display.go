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
	self.currentFocus = RUN
	charChan := make(chan string)
	runChan := make(chan string)

	if self.RunInput == nil {
		self.RunInput = inputs.Init("", [4]int{0, 20, 80, 23})
	}
	go self.RunInput.InputProcess(charChan, runChan, true)
	go manager.Run(self.Content.Manager, runChan)

	return charChan
}

func (self *Display) MkDir() chan string {
	self.currentFocus = RUN
	charChan := make(chan string)
	runChan := make(chan string)

	if self.RunInput == nil {
		self.RunInput = inputs.Init("", [4]int{0, 20, 80, 23})
	}
	go self.RunInput.InputProcess(charChan, runChan, true)
	go CreateDir(runChan, self.Content)

	return charChan
}

func (self *Display) Touch() chan string {
	self.currentFocus = RUN
	charChan := make(chan string)
	runChan := make(chan string)

	if self.RunInput == nil {
		self.RunInput = inputs.Init("", [4]int{0, 20, 80, 23})
	}
	go self.RunInput.InputProcess(charChan, runChan, true)
	go CreateFile(runChan, self.Content)

	return charChan
}

func (self *Display) Zip() chan string {
	self.currentFocus = RUN
	charChan := make(chan string)
	runChan := make(chan string)

	if self.RunInput == nil {
		self.RunInput = inputs.Init("", [4]int{0, 20, 80, 23})
	}
	go self.RunInput.InputProcess(charChan, runChan, true)
	go manager.Zipping(runChan, self.Content)

	return charChan
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
