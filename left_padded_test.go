package text

import (
	"testing"
)

func TestLeftPadMaxLine(t *testing.T) {
	cases := []struct {
		input, output  string
		maxValueLength int
		leftPad        int
	}{
		{
			"foo",
			"foo ",
			4,
			0,
		},
		{
			"foofoofoo",
			"foo…",
			4,
			0,
		},
		{
			"foo",
			"foo       ",
			10,
			0,
		},
		{
			"foo",
			"  f…",
			4,
			2,
		},
		{
			"foofoofoo",
			"  foo…",
			6,
			2,
		},
		{
			"foo",
			"  foo     ",
			10,
			2,
		},
	}

	for i, tc := range cases {
		result := LeftPadMaxLine(tc.input, tc.maxValueLength, tc.leftPad)
		if result != tc.output {
			t.Fatalf("Case %d Input:\n\n`%s`\n\nExpected Output:\n\n`%s`\n\nActual Output:\n\n`%s`",
				i, tc.input, tc.output, result)
		}
	}
}

func BenchmarkLeftPadMaxLine(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		LeftPadMaxLine("foofoofoo", 6, 2)
	}
}

func TestLeftPad(t *testing.T) {
	cases := []struct {
		input, output string
		leftPad       int
	}{
		{
			"foo",
			"foo",
			0,
		},
		{
			"foo\n",
			"foo\n",
			0,
		},
		{
			"敏捷 A quick 的狐狸 \nfox 跳过 jumps\n over a lazy 了一只懒狗 dog。",
			"    敏捷 A quick 的狐狸 \n    fox 跳过 jumps\n     over a lazy 了一只懒狗 dog。",
			4,
		},
	}

	for i, tc := range cases {
		result := LeftPad(tc.input, tc.leftPad)
		if result != tc.output {
			t.Fatalf("Case %d Input:\n\n`%s`\n\nExpected Output:\n\n`%s`\n\nActual Output:\n\n`%s`",
				i, tc.input, tc.output, result)
		}
	}
}
