package display

import (
    ui "github.com/gizak/termui"
)

func defaultHandlers(display *Display, value string) bool {
    switch value {
    case "<Up>":
        display.Content.ListUp()
    case "<Down>":
        display.Content.ListDown()
    case "<PageDown>":
        display.Content.PageDown()
    case "<PageUp>":
        display.Content.PageUp()
    case "<Enter>":
        display.Content.SelectDir()
    default:
        return false
    }

    return true
}

// func charDescript(char string, searchChan chan string) {
//     if char == "<Space>" {
//         char = " "
//     }
//     searchChan <- char
// }

// func WriteHandle(display *Display, searchChan chan string) {
//     uiEvents := ui.PollEvents()

//     for e := range uiEvents {
//         if e.Type == ui.KeyboardEvent {
//             switch e.ID {
//             case "<C-f>", "<C-q>":
//                 display.ResetList()
//                 display.ResetInput()
//                 return
//             case "<Enter>":
//                 display.SelectDir()
//                 display.ResetInput()
//                 return
//             default:
//                 if !defaultHandlers(display, e.ID) {
//                     eventID := e.ID
//                     charDescript(eventID, searchChan)
//                 }
//             }
//         }
//     }
// }

func ActionsHandle(display *Display) {
    uiEvents := ui.PollEvents()

    for {
        select {
        case e := <-uiEvents:
            switch e.ID {
            case "<C-q>":
                return
            // case "<C-f>": // Searching
            //     charChan := display.Search()
            //     WriteHandle(display, charChan)
            //     uiEvents = ui.PollEvents() // KOSTIL`
            default:
                defaultHandlers(display, e.ID)
            }
        }
    }
}
