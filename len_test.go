package text

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordLen(t *testing.T) {
	cases := []struct {
		Input  string
		Length int
	}{
		// A simple word
		{
			"foo",
			3,
		},
		// A simple word with colors
		{
			"\x1b[31mbar\x1b[0m",
			3,
		},
		// Handle prefix and suffix properly
		{
			"foo\x1b[31mfoobarHoy\x1b[0mbaaar",
			17,
		},
		// Handle chinese
		{
			"快檢什麼望對",
			12,
		},
		// Handle chinese with colors
		{
			"快\x1b[31m檢什麼\x1b[0m望對",
			12,
		},
		{
			"❌",
			2,
		},
		{
			"✔",
			2,
		},
	}

	for i, tc := range cases {
		l := WordLen(tc.Input)
		if l != tc.Length {
			t.Fatalf("Case %d Input:\n\n`%s`\n\nExpected Output:\n\n`%d`\n\nActual Output:\n\n`%d`",
				i, tc.Input, tc.Length, l)
		}
	}
}

func BenchmarkWordLen(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		WordLen("快\x1b[31m檢什麼\x1b[0m望對")
	}
}

func TestMaxLineLen(t *testing.T) {
	cases := []struct {
		text   string
		length int
	}{
		{
			`  The Lorem ipsum text is typically composed of
      pseudo-Latin words. It is commonly used as
      placeholder text to examine or demonstrate the visual
      effects of various graphic design.`,
			59,
		},
		{
			"敏捷 A \x1b31mquick\n的狐狸 fox\n跳\x1b0m过 jumps\nover a lazy\n了一只懒狗\ndog。",
			12,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.length, MaxLineLen(tc.text))
	}
}
