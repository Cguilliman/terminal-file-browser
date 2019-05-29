package main

import (
	"fmt"
	// "strings"
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

// func addCursor(value string, cursor int) string {
// 	if cursor > len(value) {
// 		for n := cursor - (len(value) - 1); n != 0; n-- {
// 			value = value + " "
// 		}
// 	}
// 	fmt.Println(value)
// 	new_string := value[:cursor] + "[" + string(value[cursor]) + "](bg:green)"
// 	if cursor < len(value) {
// 		new_string = new_string + value[cursor+1:]
// 	}
// 	return new_string
// }

func main() {
	a := map[string]interface{}{
		"s": func(a string) { fmt.Println(a) },
	}
	a["s"].(func(string))("ssss")
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
	// 		ticker = time.NewTicker(time.Second).C
	// 	case <-ticker:
	// 		drawFunction(input, "_")
	// 	}
	// }
}
