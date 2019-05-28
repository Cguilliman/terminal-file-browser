package display

import (
    "github.com/Cguilliman/terminal-file-browser/manager"
    // "github.com/Cguilliman/terminal-file-browser/inputs"
    // ui "github.com/gizak/termui"
    // "github.com/gizak/termui/widgets"
    // Input
)

type Display struct {
    Content *manager.ContentList
    // SearchInput Input
}

// func (self *Display) Search() chan string {
//     charChan := make(chan string)
//     searchChan := make(chan string)
//     go self.SearchInput.InputProcess(charChan, searchChan)
//     self.Content.SearchProcess(searchChan)
//     return charChan
// }

func InitDisplay() *Display {
    content := manager.Init("")
    // input := input.Init()

    display := &Display{
        content, 
        // input, 
    }
    return display
}
