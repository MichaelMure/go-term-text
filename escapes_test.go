package text

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractApplyTermEscapes(t *testing.T) {
	cases := []struct {
		Name        string
		Input       string
		Output      string
		TermEscapes []EscapeItem
	}{

		{
			"A plain ascii line with escapes",
			"This \x1b[31mis an\x1b[0m example.",
			"This is an example.",
			[]EscapeItem{{"\x1b[31m", 5}, {"\x1b[0m", 10}},
		},

		{
			"Escape at the end",
			"This \x1b[31mis an example.\x1b[0m",
			"This is an example.",
			[]EscapeItem{{"\x1b[31m", 5}, {"\x1b[0m", 19}},
		},

		{
			"A plain wide line with escapes",
			"一只敏捷\x1b[31m的狐狸\x1b[0m跳过了一只懒狗。",
			"一只敏捷的狐狸跳过了一只懒狗。",
			[]EscapeItem{{"\x1b[31m", 4}, {"\x1b[0m", 7}},
		},

		{
			"A normal-wide mixed line with escapes",
			"一只 A Quick 敏捷\x1b[31m的狐 Fox 狸\x1b[0m跳过了Dog一只懒狗。",
			"一只 A Quick 敏捷的狐 Fox 狸跳过了Dog一只懒狗。",
			[]EscapeItem{{"\x1b[31m", 13}, {"\x1b[0m", 21}},
		},

		{
			"Multiple escapes at the same place",
			"\x1b[1m\x1b[31mThis \x1b[1m\x1b[31mis an\x1b[0m example.\x1b[1m\x1b[31m",
			"This is an example.",
			[]EscapeItem{
				{"\x1b[1m", 0}, {"\x1b[31m", 0},
				{"\x1b[1m", 5}, {"\x1b[31m", 5},
				{"\x1b[0m", 10},
				{"\x1b[1m", 19}, {"\x1b[31m", 19}},
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			cleaned, escapes := ExtractTermEscapes(tc.Input)

			assert.Equal(t, tc.Output, cleaned)
			assert.Equal(t, tc.TermEscapes, escapes)

			augmented := ApplyTermEscapes(cleaned, escapes)

			assert.Equal(t, tc.Input, augmented)
		})
	}
}

func BenchmarkExtractTermEscapes(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		ExtractTermEscapes("\x1b[1m\x1b[31mThis \x1b[1m\x1b[31mis an\x1b[0m example.\x1b[1m\x1b[31m")
	}
}

func BenchmarkApplyTermEscapes(b *testing.B) {
	cleaned, escapes := ExtractTermEscapes("\x1b[1m\x1b[31mThis \x1b[1m\x1b[31mis an\x1b[0m example.\x1b[1m\x1b[31m")

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		ApplyTermEscapes(cleaned, escapes)
	}
}
