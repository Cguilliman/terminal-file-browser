package inputs

const STYLE string = "](bg:green)"

type Cursor struct {
    cursor int
    ping bool
    value string
}

func (self *Cursor) addCursor() string {
	if cursor > len(value) {
        // append dynamic spaces amount
		for n := cursor - (len(value) - 1); n != 0; n-- {
			value = value + " "
		}
	}
	new_string := value[:cursor] + "[" + string(value[cursor]) + STYLE
	if cursor < len(value) {
		new_string = new_string + value[cursor+1:]
	}
	return new_string
}

func (self *Cursor) removeCursor() string {
    return self.value[:self.cursor] + self.value[self.cursor+1] + self.value[self.cursor+len(STYLE)+1]
}

func (self *Cursor) leftCursor() {
    value = self.removeCursor()
    self.cursor--
    // return 
}

func (self *Cursor) rightCursor() {
    self.cursor++
}
