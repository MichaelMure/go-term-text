package text

import (
	"strings"
	"unicode/utf8"
)

// EscapeItem hold the description of terminal escapes in a line.
// 'item' is the actual escape command
// 'pos' is the index in the rune array where the 'item' shall be inserted back.
// For example, the escape item in "F\x1b33mox" is {"\x1b33m", 1}.
type EscapeItem struct {
	Item string
	Pos  int
}

// ExtractTermEscapes extract terminal escapes out of a line and returns a new
// line without terminal escapes and a slice of escape items. The terminal escapes
// can be inserted back into the new line at rune index 'item.pos' to recover the
// original line.
//
// Required: The line shall not contain "\n"
func ExtractTermEscapes(line string) (string, []EscapeItem) {
	var termEscapes []EscapeItem
	var line1 strings.Builder

	pos := 0
	item := ""
	occupiedRuneCount := 0
	inEscape := false
	for i, r := range []rune(line) {
		if r == '\x1b' {
			pos = i
			item = string(r)
			inEscape = true
			continue
		}
		if inEscape {
			item += string(r)
			if r == 'm' {
				termEscapes = append(termEscapes, EscapeItem{item, pos - occupiedRuneCount})
				occupiedRuneCount += utf8.RuneCountInString(item)
				inEscape = false
			}
			continue
		}
		line1.WriteRune(r)
	}

	return line1.String(), termEscapes
}

// ApplyTermEscapes apply the extracted terminal escapes to the edited line.
// The only edit allowed is to insert "\n" like that in softwrapLine.
// Callers shall ensure this since this function is not able to check it.
func ApplyTermEscapes(line string, escapes []EscapeItem) string {
	if len(escapes) == 0 {
		return line
	}

	var out strings.Builder

	currPos := 0
	currItem := 0
	for _, r := range line {
		for currItem < len(escapes) && currPos == escapes[currItem].Pos {
			out.WriteString(escapes[currItem].Item)
			currItem++
		}
		out.WriteRune(r)
		currPos++
	}

	// Don't forget the trailing escapes, if any.
	for currItem < len(escapes) && currPos >= escapes[currItem].Pos {
		out.WriteString(escapes[currItem].Item)
		currItem++
	}

	return out.String()
}
