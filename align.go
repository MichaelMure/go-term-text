package text

import (
	"strings"
	"unicode"
)

type Alignment int

const (
	NoAlign Alignment = iota
	AlignLeft
	AlignCenter
	AlignRight
)

// LineAlign align the given line as asked and apply the needed padding to match the given
// lineWidth, while ignoring the terminal escape sequences.
// If the given lineWidth is too small to fit the given line, it's returned without
// padding, overflowing lineWidth.
func LineAlign(line string, lineWidth int, align Alignment) string {
	switch align {
	case NoAlign:
		return line
	case AlignLeft:
		return LineAlignLeft(line, lineWidth)
	case AlignCenter:
		return LineAlignCenter(line, lineWidth)
	case AlignRight:
		return LineAlignRight(line, lineWidth)
	}
	panic("unknown alignment")
}

// LineAlignLeft align the given line on the left while ignoring the terminal escape sequences.
// If the given lineWidth is too small to fit the given line, it's returned without
// padding, overflowing lineWidth.
func LineAlignLeft(line string, lineWidth int) string {
	cleaned, escapes := ExtractTermEscapes(line)
	trimmed := strings.TrimLeftFunc(cleaned, unicode.IsSpace)
	recomposed := ApplyTermEscapes(trimmed, escapes)
	return recomposed
}

// LineAlignCenter align the given line on the center and apply the needed left
// padding, while ignoring the terminal escape sequences.
// If the given lineWidth is too small to fit the given line, it's returned without
// padding, overflowing lineWidth.
func LineAlignCenter(line string, lineWidth int) string {
	cleaned, escapes := ExtractTermEscapes(line)
	trimmed := strings.TrimFunc(cleaned, unicode.IsSpace)
	recomposed := ApplyTermEscapes(trimmed, escapes)
	totalPadLen := lineWidth - WordLen(trimmed)
	if totalPadLen < 0 {
		totalPadLen = 0
	}
	pad := strings.Repeat(" ", totalPadLen/2)
	return pad + recomposed
}

// LineAlignRight align the given line on the right and apply the needed left
// padding to match the given lineWidth, while ignoring the terminal escape sequences.
// If the given lineWidth is too small to fit the given line, it's returned without
// padding, overflowing lineWidth.
func LineAlignRight(line string, lineWidth int) string {
	cleaned, escapes := ExtractTermEscapes(line)
	trimmed := strings.TrimRightFunc(cleaned, unicode.IsSpace)
	recomposed := ApplyTermEscapes(trimmed, escapes)
	padLen := lineWidth - WordLen(trimmed)
	if padLen < 0 {
		padLen = 0
	}
	pad := strings.Repeat(" ", padLen)
	return pad + recomposed
}
