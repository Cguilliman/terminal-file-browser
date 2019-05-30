package inputs

import (
	"strings"
)

const DEFAULT_MESSAGE string = "Input field"

type Input struct {
	RenderChan chan string
	Widget     *Widget
	Value      string
}

func (self *Input) Reset() {
	self.RenderChan <- DEFAULT_MESSAGE
}

func (self *Input) InputProcess(charChan, responseChan chan string, isClose bool) {
	self.Value = ""
	cursor := DefaultCursor(
		self.RenderChan,
		self.Value,
	)
	if isClose {
		defer close(responseChan)
	}
	defer cursor.Disable()
	defer self.Reset()

	for char := range charChan {
		switch {
		case char == "<Backspace>" && cursor.Index > 0:
			if cursor.Index < len(self.Value) {
				self.Value = strings.Join([]string{
					self.Value[:cursor.Index-1],
					self.Value[cursor.Index:],
				}, "")
			} else {
				self.Value = self.Value[:len(self.Value)-1]
			}

			cursor.LeftCursor(self.Value)
		case char == "<C-<Backspace>>" && cursor.Index > 0:
			splited := strings.Split(self.Value[:cursor.Index-1], " ")
			newValue := strings.Join(splited[:len(splited)-1], " ")
			// mb add space before deleted pattern
			if cursor.Index < len(self.Value) {
				self.Value = newValue + self.Value[cursor.Index:]
			} else {
				self.Value = newValue
			}

			cursor.Move(
				len(newValue),
				self.Value,
			)
		case char == "<Delete>" && cursor.Index < len(self.Value):
			if cursor.Index > 0 {
				self.Value = strings.Join([]string{
					self.Value[:cursor.Index],
					self.Value[cursor.Index+1:],
				}, "")
			} else if len(self.Value) > 1 {
				self.Value = self.Value[cursor.Index+1:]
			} else {
				self.Value = ""
			}

			cursor.Move(
				cursor.Index,
				self.Value,
			)
		case char == "<Left>" && cursor.Index > 0:
			cursor.LeftCursor(self.Value)
		case char == "<Right>" && cursor.Index < len(self.Value):
			cursor.RightCursor(self.Value)
		case len(char) > 1:
			continue
		default:
			self.Value = strings.Join([]string{
				self.Value[:cursor.Index],
				char,
				self.Value[cursor.Index:],
			}, "")

			cursor.RightCursor(self.Value)
		}

		responseChan <- self.Value // will return in parent channel of value
	}
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
