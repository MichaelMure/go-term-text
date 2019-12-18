package fuzzing

import (
	"encoding/binary"
	"fmt"
	"strings"

	text "github.com/MichaelMure/go-term-text"
)

func Fuzz(data []byte) int {
	if len(data) < 8 {
		return -1
	}

	lineWidth := int(binary.BigEndian.Uint16(data[:2]))
	data = data[2:]

	indentln := int(binary.BigEndian.Uint16(data[:2]))
	data = data[2:]
	indent := strings.Repeat(" ", indentln)

	padln := int(binary.BigEndian.Uint16(data[:2]))
	data = data[2:]
	pad := strings.Repeat(" ", padln)

	align := text.Alignment(binary.BigEndian.Uint16(data[:2]))
	data = data[2:]

	t := string(data)

	fmt.Println(lineWidth, indentln, padln, align)

	text.WrapWithPadIndentAlign(t, lineWidth, indent, pad, align)
	return 1
}
