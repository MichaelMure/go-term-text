package text

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLineAlignLeft(t *testing.T) {
	cases := []struct {
		line   string
		width  int
		output string
	}{
		{
			"  foo foo bar",
			30,
			"foo foo bar",
		},
		{
			"The Lorem ipsum text is typically composed of pseudo-Latin words.",
			70,
			"The Lorem ipsum text is typically composed of pseudo-Latin words.",
		},
		// width too low return the same input
		{
			"The Lorem ipsum text is typically composed of pseudo-Latin words.",
			10,
			"The Lorem ipsum text is typically composed of pseudo-Latin words.",
		},
		// respect escape sequences and wide chars
		{
			"   敏捷 A \x1b31mquick\n的狐狸 fox\n跳\x1b0m过 jumps\nover a lazy\n了一只懒狗\ndog。",
			60,
			"敏捷 A \x1b31mquick\n的狐狸 fox\n跳\x1b0m过 jumps\nover a lazy\n了一只懒狗\ndog。",
		},
	}
	for _, tc := range cases {
		out := LineAlignLeft(tc.line, tc.width)
		assert.Equal(t, tc.output, out)
	}
}

func BenchmarkLineAlignLeft(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		LineAlignLeft("敏捷 A \x1b31mquick\n的狐狸 fox\n跳\x1b0m过 jumps\nover a lazy\n了一只懒狗\ndog。", 60)
	}
}

func TestLineAlignCenter(t *testing.T) {
	cases := []struct {
		line   string
		width  int
		output string
	}{
		{
			"The Lorem ipsum text is typically composed of pseudo-Latin words.",
			75,
			"     The Lorem ipsum text is typically composed of pseudo-Latin words.",
		},
		// width too low return the same input
		{
			"The Lorem ipsum text is typically composed of pseudo-Latin words.",
			10,
			"The Lorem ipsum text is typically composed of pseudo-Latin words.",
		},
		// respect escape sequences and wide chars
		{
			"敏捷 A \x1b31mquick\n的狐狸 fox\n跳\x1b0m过 jumps\nover a lazy\n了一只懒狗\ndog。",
			60,
			" 敏捷 A \x1b31mquick\n的狐狸 fox\n跳\x1b0m过 jumps\nover a lazy\n了一只懒狗\ndog。",
		},
	}
	for _, tc := range cases {
		out := LineAlignCenter(tc.line, tc.width)
		assert.Equal(t, tc.output, out)
	}
}

func BenchmarkLineAlignCenter(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		LineAlignCenter("敏捷 A \x1b31mquick\n的狐狸 fox\n跳\x1b0m过 jumps\nover a lazy\n了一只懒狗\ndog。", 60)
	}
}

func TestLineAlignRight(t *testing.T) {
	cases := []struct {
		line   string
		width  int
		output string
	}{
		{
			"The Lorem ipsum text is typically composed of pseudo-Latin words.",
			70,
			"     The Lorem ipsum text is typically composed of pseudo-Latin words.",
		},
		// width too low return the same input
		{
			"The Lorem ipsum text is typically composed of pseudo-Latin words.",
			10,
			"The Lorem ipsum text is typically composed of pseudo-Latin words.",
		},
		// respect escape sequences and wide chars
		{
			"敏捷 A \x1b31mquick\n的狐狸 fox\n跳\x1b0m过 jumps\nover a lazy\n了一只懒狗\ndog。",
			60,
			"  敏捷 A \x1b31mquick\n的狐狸 fox\n跳\x1b0m过 jumps\nover a lazy\n了一只懒狗\ndog。",
		},
	}
	for _, tc := range cases {
		out := LineAlignRight(tc.line, tc.width)
		assert.Equal(t, tc.output, out)
	}
}

func BenchmarkLineAlignRight(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		LineAlignRight("敏捷 A \x1b31mquick\n的狐狸 fox\n跳\x1b0m过 jumps\nover a lazy\n了一只懒狗\ndog。", 60)
	}
}
