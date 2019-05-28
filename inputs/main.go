package inputs

import (
	"strings"
	// ui "github.com/gizak/termui"
	// "github.com/gizak/termui/widgets"
)

const (
	DEFAULT_MESSAGE string = "Input field"
	// PINGIN_CHAR string = "|"
)

type Input struct {
	RenderChan chan string
	Widget     *Widget
	Value      string
}

func (self *Input) Reset() {
	self.RenderChan <- DEFAULT_MESSAGE
}

func Init(value string, size [4]int) *Input {
	var input Input
	input.RenderChan, input.Widget = initWidget(size)

	if value == "" {
		input.Value = DEFAULT_MESSAGE
		input.Reset()
	} else {
		input.Value = value
		input.RenderChan <- value
	}

	return &input
}

func (self *Input) InputProcess(charChan, responseChan chan string) {
	var value string
	self.Value = ""
	// var cursor int = len(value)
	// [color](fg:white,bg:green) output -- ping style
	ticker := time.NewTicker(time.Second).C

	for {
		select {
		case char := <-charChan:
			switch {
			case char == "<Backspace>" && len(value) > 0:
				value = value[:len(value)-1]
			case char == "<C-<Backspace>>" && len(value) > 0:
				splited := strings.Split(value, " ")
				value = strings.Join(splited[:len(splited)-1], " ")
			// case char == "<Left>" && cursor > 0:
			//     cursor--
			// case char == "<Right>" && cursor < len(value):
			//     cursor++
			case len(char) > 1:
				continue
			default:
				value = value + char
			}

			self.Value = value       // set in local variable
			responseChan <- value    // will return in parent channel of value
			self.RenderChan <- value // will render string in input

			ticker = time.NewTicker(time.Second).C
		case <-ticker:
			// TODO
		}
	}

	// for char := range charChan {
	//     switch {
	//     case char == "<Backspace>" && len(value) > 0:
	//         value = value[:len(value)-1]
	//     case char == "<C-<Backspace>>" && len(value) > 0:
	//         splited := strings.Split(value, " ")
	//         value = strings.Join(splited[:len(splited)-1], " ")
	//     // case char == "<Left>" && cursor > 0:
	//     //     cursor--
	//     // case char == "<Right>" && cursor < len(value):
	//     //     cursor++
	//     case len(char) > 1:
	//         continue
	//     default:
	//         value = value + char
	//     }

	//     self.Value = value  // set in local variable
	//     responseChan <- value  // will return in parent channel of value
	//     self.RenderChan <- value // will render string in input
	// }
}
