package display

import (
    "github.com/Cguilliman/terminal-file-browser/manager"
    "github.com/Cguilliman/terminal-file-browser/inputs"
)

type Display struct {
    Content *manager.ContentList
    SearchInput *inputs.Input
}

func (self *Display) Search() chan string {
    charChan := make(chan string)
    searchChan := make(chan string)

    go self.SearchInput.InputProcess(charChan, searchChan)
    go self.Content.SearchProcess(searchChan)

    return charChan
}

func InitDisplay() *Display {
    content := manager.Init("")
    saerchInput := inputs.Init("", [4]int{0, 0, 80, 3})

    display := &Display{
        content, 
        saerchInput, 
    }
    return display
}
