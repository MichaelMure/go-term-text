package text

import (
	"strings"

	"github.com/mattn/go-runewidth"
)

// Len return the length of a string in a terminal, while ignoring the terminal
// escape sequences.
func Len(text string) int {
	length := 0
	var ed EscapeDetector

	for _, r := range []rune(text) {
		ed.Witness(r)
		if !ed.InEscape() {
			length += runewidth.RuneWidth(r)
		}
	}

	return length
}

// MaxLineLen return the length in a terminal of the longest line, while
// ignoring the terminal escape sequences.
func MaxLineLen(text string) int {
	lines := strings.Split(text, "\n")

	max := 0

	for _, line := range lines {
		length := Len(line)
		if length > max {
			max = length
		}
	}

	return max
}
