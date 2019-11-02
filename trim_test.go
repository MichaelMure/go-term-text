package text

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrimSpace(t *testing.T) {
	cases := []struct {
		line   string
		output string
	}{
		{
			"foo",
			"foo",
		},
		{
			"foo      ",
			"foo",
		},
		{
			"      foo",
			"foo",
		},
		{
			"   \x1b[31mbar\x1b[0m     ",
			"\x1b[31mbar\x1b[0m",
		},
		{
			"\x1b[31m   bar     \x1b[0m",
			"\x1b[31mbar\x1b[0m",
		},
		{
			"  \x1b[31m   bar     \x1b[0m   ",
			"\x1b[31mbar\x1b[0m",
		},
		{
			"  敏捷 A \x1b31mquick\n的狐狸 fox\n跳\x1b0m过 jumps\nover a lazy\n了一只懒狗\ndog。   ",
			"敏捷 A \x1b31mquick\n的狐狸 fox\n跳\x1b0m过 jumps\nover a lazy\n了一只懒狗\ndog。",
		},
	}
	for _, tc := range cases {
		out := TrimSpace(tc.line)
		assert.Equal(t, tc.output, out)
	}
}

func BenchmarkTrimSpace(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		TrimSpace("  敏捷 A \x1b31mquick\n的狐狸 fox\n跳\x1b0m过 jumps\nover a lazy\n了一只懒狗\ndog。   ")
	}
}
