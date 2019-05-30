package display

import (
    "github.com/Cguilliman/terminal-file-browser/inputs"
    "github.com/Cguilliman/terminal-file-browser/manager"
)

func Command(dispaly *Display, command func(chan string, *manager.ContentList)) chan string {
    dispaly.currentFocus = RUN
    charChan := make(chan string)
    runChan := make(chan string)

    if dispaly.RunInput == nil {
        dispaly.RunInput = inputs.Init("", [4]int{0, 20, 80, 23})
    }
    go dispaly.RunInput.InputProcess(charChan, runChan, true)
    go command(runChan, dispaly.Content)

    return charChan
}
