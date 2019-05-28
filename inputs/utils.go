package inputs

const STYLE string = "](bg:green)"

func addCursor(value string, cursor int) string {
	if cursor > len(value) {
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

func removeCursor(value string, cursor int) string {

}
