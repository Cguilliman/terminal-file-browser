package display

import (
    mg "github.com/Cguilliman/terminal-file-browser/manager"
    "github.com/Cguilliman/terminal-file-browser/input"
    ui "github.com/gizak/termui"
    "github.com/gizak/termui/widgets"
    // Input
)

type Display type {
    Content *mg.ContentList
    // SearchInput Input
}

func (self *Display) Search() chan string {
    charChan := make(chan string)
    searchChan := make(chan string)
    go self.SearchInput.InputProcess(charChan, searchChan)
    self.Content.SearchProcess(searchChan)
    return charChan
}

func InitDisplay() *Display {
    content := mg.Init()
    input := input.Init()

    display := &Display{
        content, 
        input, 
    }
    return display
}
