package display

import (
	"strings"

	ui "github.com/gizak/termui"
	"github.com/gizak/termui/widgets"
)

const (
	DEFAULT_MESSAGE string = "Input field"
	// PINGIN_CHAR string = "|"
)

type Input struct {
	Widget *widgets.Paragraph
	Data   string
}

func (self *Input) InputProcess(charChan, searchChan chan string) {
	var searchString string
	// var cursor int = len(searchString)

	for char := range charChan {
		switch {
		case char == "<Backspace>" && len(searchString) > 0:
			searchString = searchString[:len(searchString)-1]
		case char == "<C-<Backspace>>" && len(searchString) > 0:
			splitedSearch := strings.Split(searchString, " ")
			searchString = strings.Join(splitedSearch[:len(splitedSearch)-1], " ")
		// case char == "<Left>" && cursor > 0:
		//     cursor--
		// case char == "<Right>" && cursor < len(searchString):
		//     cursor++
		case len(char) > 1:
			continue
		default:
			searchString = searchString + char
		}

		searchChan <- searchString
		self.InputText(searchString, true)
	}
}

func (self *Input) InputText(value string, isRerender bool) {
	self.Data = value
	self.Reset(isRerender)
}

func (self *Input) Reset(isRerender bool) {
	if isRerender {
		defer ui.Render(self.Widget)
	}
	self.Widget.Text = self.Data
}

func InitInput(isRerender bool) *Input {
	input := Input{
		widgets.NewParagraph(),
		DEFAULT_MESSAGE,
	}
	input.Widget.SetRect(0, 0, 80, 3)
	input.Reset(isRerender)
	return &input
}
