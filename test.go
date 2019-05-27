package main

import (
	"fmt"
	"strings"
	// ui "github.com/gizak/termui"
	// "github.com/gizak/termui/widgets"
	// "time"
)

// func drawFunction(input *widgets.Paragraph, new_char string) {
// 	if new_char == "" {

// 	}
// 	input.Text = input.Text + new_char
// 	ui.Render(input)
// }

func main() {
	fmt.Println(strings.Contains("ss", "aass"))
	// if err := ui.Init(); err != nil {
	// 	panic(err)
	// }

	// input := widgets.NewParagraph()
	// input.Text = "Search Field _"
	// input.SetRect(0, 0, 80, 3)
	// ui.Render(input)

	// defer ui.Close()

	// uiEvents := ui.PollEvents()
	// ticker := time.NewTicker(time.Second).C
	// for {
	// 	select {
	// 	case e := <-uiEvents:
	// 		switch e.ID { // event string/identifier
	// 		case "<C-q>", "<C-c>": // press 'q' or 'C-c' to quit
	// 			return
	// 		}
	// 		switch e.Type {
	// 		case ui.KeyboardEvent: // handle all key presses
	// 			eventID := e.ID // keypress string
	// 			if eventID == "<Space>" {
	// 				eventID = " "
	// 			}
	// 			drawFunction(input, eventID)
	// 		}
	// 	case <-ticker:
	// 		drawFunction(input, "")
	// 	}
	// }
}
