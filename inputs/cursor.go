package inputs

import (
	"strings"
	"time"
	// "fmt"
)

const STYLE string = "](bg:green)"

type Cursor struct {
	Index       int
    pingingChan chan string
    renderChan  chan string
	isPing      bool
	blankValue  string
}

// will add cursor style and send value in render channel
func (self *Cursor) AddCursor(isSend bool) string {
	if self.Index >= len(self.blankValue) {
		// append dynamic spaces amount
		for n := self.Index - (len(self.blankValue) - 1); n != 0; n-- {
			self.blankValue = self.blankValue + " "
		}
	}

	value := strings.Join([]string{
		self.blankValue[:self.Index],
		"[",
		string(self.blankValue[self.Index]),
		STYLE,
	}, "")

	if self.Index < len(self.blankValue) {
		value = value + self.blankValue[self.Index+1:]
	}
	if isSend {
		self.renderChan <- value
	}
	return value
}

// will remove cursor style and send value in render channel
func (self *Cursor) RemoveCursor(isSend bool) string {
	if isSend {
		self.renderChan <- self.blankValue
	}
	return self.blankValue
}

// move cursor left
// add it in value and send in render channel
func (self *Cursor) LeftCursor(newValue string) {
	// add guard
	self.Move(self.Index-1, newValue)
}

// move cursor right
// add it in value and send in render channel
func (self *Cursor) RightCursor(newValue string) {
	// add guard
    self.Move(self.Index+1, newValue)
}

func (self *Cursor) Move(index int, newValue string) {
    self.Index = index
    self.pingingChan <- newValue
}

func (self *Cursor) ping() {
	if self.isPing {
		self.RemoveCursor(true)
		self.isPing = false
	} else {
		self.AddCursor(true)
		self.isPing = true
	}
}

// pinging process
func (self *Cursor) Pinging() {
	ticker := time.NewTicker(time.Second / 2).C
	for {
		select {
		case value, ok := <-self.pingingChan:
			self.blankValue = value
			self.AddCursor(true)
			ticker = time.NewTicker(time.Second / 2).C
			if !ok {
				return
			}
		case <-ticker:
			self.ping()
		}
	}
}

func (self *Cursor) Disable() {
	close(self.pingingChan)
}

// initialize default cursor
// start pinging
func DefaultCursor(renderChan chan string, value string) *Cursor {
	pingingChan := make(chan string)
	obj := Cursor{
		Index:       len(value),
        renderChan:  renderChan,
        pingingChan: pingingChan,
		isPing:      false,
		blankValue:  value,
	}
	go obj.Pinging()
	return &obj
}
