package text

import (
	"strings"
	"unicode"
)

// TrimSpace remove the leading and trailing whitespace while ignoring the
// terminal escape sequences.
// Returns the trimmed
func TrimSpace(line string) (result string, left int, right int) {
	cleaned, escapes := ExtractTermEscapes(line)

	// trim left while counting
	trimmed := strings.TrimLeftFunc(cleaned, func(r rune) bool {
		if unicode.IsSpace(r) {
			left++
			return true
		}
		return false
	})

	// trim right while counting
	trimmed = strings.TrimRightFunc(trimmed, func(r rune) bool {
		if unicode.IsSpace(r) {
			right++
			return true
		}
		return false
	})

	// offset the escape sequences, bounded in the trimmed string space
	for i, seq := range escapes {
		if seq.Pos-left < 0 {
			escapes[i].Pos = 0
		} else if seq.Pos-left > len(trimmed) {
			escapes[i].Pos = len(trimmed)
		} else {
			escapes[i].Pos = seq.Pos - left
		}
	}

	result = ApplyTermEscapes(trimmed, escapes)
	return
}
