package text

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestApplyTermEscapes(t *testing.T) {
	cases := []struct {
		Name        string
		Input       string
		Output      string
		TermEscapes []EscapeItem
	}{
		{
			"negative offset",
			"This is an example.",
			"\x1b[31mThis is an\x1b[0m example.",
			[]EscapeItem{{"\x1b[31m", -5}, {"\x1b[0m", 10}},
		},
		{
			"offset too far",
			"This is an example.",
			"This \x1b[31mis an example.\x1b[0m",
			[]EscapeItem{{"\x1b[31m", 5}, {"\x1b[0m", 30}},
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			result := ApplyTermEscapes(tc.Input, tc.TermEscapes)
			assert.Equal(t, tc.Output, result)
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

func TestOffsetEscapes(t *testing.T) {
	cases := []struct {
		input  []EscapeItem
		offset int
		output []EscapeItem
	}{
		{
			[]EscapeItem{{Pos: 0}, {Pos: 2}, {Pos: 20}},
			5,
			[]EscapeItem{{Pos: 5}, {Pos: 7}, {Pos: 25}},
		},
		{
			[]EscapeItem{{Pos: 0}, {Pos: 2}, {Pos: 20}},
			-5,
			[]EscapeItem{{Pos: -5}, {Pos: -3}, {Pos: 15}},
		},
	}
	for _, tc := range cases {
		result := OffsetEscapes(tc.input, tc.offset)
		assert.Equal(t, tc.output, result)
	}
}

func TestEscapeDetector(t *testing.T) {
	input := "This \u001B[31mis an\u001B[0m example."
	states := []struct {
		inEscape bool
		started  bool
		ended    bool
	}{
		{false, false, false}, // T
		{false, false, false}, // h
		{false, false, false}, // i
		{false, false, false}, // s
		{false, false, false}, //
		{true, true, false},   // \u001b
		{true, false, false},  // [
		{true, false, false},  // 3
		{true, false, false},  // 1
		{true, false, true},   // m
		{false, false, false}, // i
		{false, false, false}, // s
		{false, false, false}, //
		{false, false, false}, // a
		{false, false, false}, // n
		{true, true, false},   // \u001b
		{true, false, false},  // [
		{true, false, false},  // 0
		{true, false, true},   // m
		{false, false, false}, //
		{false, false, false}, // e
		{false, false, false}, // x
		{false, false, false}, // a
		{false, false, false}, // m
		{false, false, false}, // p
		{false, false, false}, // l
		{false, false, false}, // e
		{false, false, false}, // .
	}

	var ed EscapeDetector
	for i, r := range input {
		ed.Witness(r)
		require.Equal(t, states[i].inEscape, ed.InEscape())
		require.Equal(t, states[i].started, ed.Started())
		require.Equal(t, states[i].ended, ed.Ended())
	}
}
