package text

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrimSpace(t *testing.T) {
	cases := []struct {
		line        string
		output      string
		left, right int
	}{
		{
			"foo",
			"foo",
			0, 0,
		},
		{
			"foo      ",
			"foo",
			0, 6,
		},
		{
			"      foo",
			"foo",
			6, 0,
		},
		{
			"   \x1b[31mbar\x1b[0m     ",
			"\x1b[31mbar\x1b[0m",
			3, 5,
		},
		{
			"\x1b[31m   bar     \x1b[0m",
			"\x1b[31mbar\x1b[0m",
			3, 5,
		},
		{
			"  \x1b[31m   bar     \x1b[0m   ",
			"\x1b[31mbar\x1b[0m",
			5, 8,
		},
		{
			"  敏捷 A \x1b31mquick\n的狐狸 fox\n跳\x1b0m过 jumps\nover a lazy\n了一只懒狗\ndog。   ",
			"敏捷 A \x1b31mquick\n的狐狸 fox\n跳\x1b0m过 jumps\nover a lazy\n了一只懒狗\ndog。",
			2, 3,
		},
	}
	for _, tc := range cases {
		out, left, right := TrimSpace(tc.line)
		assert.Equal(t, tc.output, out)
		assert.Equal(t, tc.left, left)
		assert.Equal(t, tc.right, right)
	}
}

func BenchmarkTrimSpace(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		TrimSpace("  敏捷 A \x1b31mquick\n的狐狸 fox\n跳\x1b0m过 jumps\nover a lazy\n了一只懒狗\ndog。   ")
	}
}
