package display

import (
	"github.com/Cguilliman/terminal-file-browser/inputs"
	"github.com/Cguilliman/terminal-file-browser/manager"
)

type Display struct {
	Content      *manager.ContentList
	SearchInput  *inputs.Input
	RunInput     *inputs.Input
	currentFocus string
}

func (self *Display) Submit(charChan chan string) bool {
	// TODO re-factor
	if self.currentFocus == "search" {
		self.Content.SelectDir()
		// self.SearchInput.Reset()
		close(charChan)
		return true
	} else if self.currentFocus == "run" {
		// self.Content.SelectDir()
		// self.RunInput.Reset()
		close(charChan)
		return true
	}
	return false
}

func (self *Display) Search() chan string {
	self.currentFocus = "search"
	charChan := make(chan string)
	searchChan := make(chan string)

	go self.SearchInput.InputProcess(charChan, searchChan)
	go self.Content.SearchProcess(searchChan)

	return charChan
}

func (self *Display) Run() chan string {
	self.currentFocus = "run"
	charChan := make(chan string)
	runChan := make(chan string)

	self.RunInput = inputs.Init("", [4]int{0, 3, 80, 3})
	go self.RunInput.InputProcess(charChan, runChan)

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
