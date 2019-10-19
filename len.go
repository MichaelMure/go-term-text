package text

import (
	"strings"

	"github.com/mattn/go-runewidth"
)

// WordLen return the length of a word, while ignoring the terminal escape
// sequences
func WordLen(word string) int {
	length := 0
	escape := false

	for _, char := range word {
		if char == '\x1b' {
			escape = true
		}
		if !escape {
			length += runewidth.RuneWidth(rune(char))
		}
		if char == 'm' {
			escape = false
		}
	}

	return length
}

// MaxLineLen return the length of the longest line, while ignoring the terminal escape sequences
func MaxLineLen(text string) int {
	lines := strings.Split(text, "\n")

	max := 0

	for _, line := range lines {
		length := WordLen(line)
		if length > max {
			max = length
		}
	}

	return max
}
