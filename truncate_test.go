package text

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTruncateMax(t *testing.T) {
	cases := []struct {
		line   string
		output string
		length int
	}{
		{
			"foo",
			"foo",
			6,
		},
		{
			"foobarfoobar",
			"fooba…",
			6,
		},
		{
			"foo",
			"…",
			0,
		},
		{
			"foo",
			"…",
			1,
		},
		{
			"\x1b[31mbar\x1b[0m",
			"\x1b[31mbar\x1b[0m",
			3,
		},
		{
			"\x1b[31mbar\x1b[0m",
			"\x1b[31mb\x1b[0m…",
			2,
		},
		{
			"敏捷 A \x1b31mquick 的狐狸 fox 跳\x1b0m过 jumps over a lazy 了一只懒狗 dog。",
			"敏捷 A \x1b31mquick \x1b0m…",
			15,
		},
	}
	for _, tc := range cases {
		out := TruncateMax(tc.line, tc.length)
		assert.Equal(t, tc.output, out)
	}
}
