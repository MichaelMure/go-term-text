package text

import (
	"reflect"
	"strings"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestWrap(t *testing.T) {
	cases := []struct {
		Input, Output string
		Lim           int
	}{
		// A simple word passes through.
		{
			"foo",
			"foo",
			4,
		},
		// Word breaking
		{
			"foobarbaz",
			"foob\narba\nz",
			4,
		},
		// Lines are broken at whitespace.
		{
			"foo bar baz",
			"foo\nbar\nbaz",
			4,
		},
		// Word breaking
		{
			"foo bars bazzes",
			"foo\nbars\nbazz\nes",
			4,
		},
		// A word that would run beyond the width is wrapped.
		{
			"fo sop",
			"fo\nsop",
			4,
		},
		// A tab counts as 4 characters.
		{
			"foo\nb\t r\n baz",
			"foo\nb\nr\n baz",
			4,
		},
		// Trailing whitespace is removed after used for wrapping.
		// Runs of whitespace on which a line is broken are removed.
		{
			"foo    \nb   ar   ",
			"foo\n\nb\nar\n",
			4,
		},
		// An explicit line break at the end of the input is preserved.
		{
			"foo bar baz\n",
			"foo\nbar\nbaz\n",
			4,
		},
		// Explicit break are always preserved.
		{
			"\nfoo bar\n\n\nbaz\n",
			"\nfoo\nbar\n\n\nbaz\n",
			4,
		},
		// Ignore complete words with terminal color sequence
		{
			"foo \x1b[31mbar\x1b[0m baz",
			"foo\n\x1b[31mbar\x1b[0m\nbaz",
			4,
		},
		// Handle words with colors sequence inside the word
		{
			"foo b\x1b[31mbar\x1b[0mr baz",
			"foo\nb\x1b[31mbar\n\x1b[0mr\nbaz",
			4,
		},
		// Break words with colors sequence inside the word
		{
			"foo bb\x1b[31mbar\x1b[0mr baz",
			"foo\nbb\x1b[31mba\nr\x1b[0mr\nbaz",
			4,
		},
		// Complete example:
		{
			" This is a list: \n\n\t* foo\n\t* bar\n\n\n\t* baz  \nBAM    ",
			" This\nis a\nlist:\n\n    *\nfoo\n    *\nbar\n\n\n    *\nbaz\nBAM\n",
			6,
		},
		// Handle chinese (wide characters)
		{
			"一只敏捷的狐狸跳过了一只懒狗。",
			"一只敏捷的狐\n狸跳过了一只\n懒狗。",
			12,
		},
		// Handle chinese with colors
		{
			"一只敏捷的\x1b[31m狐狸跳过\x1b[0m了一只懒狗。",
			"一只敏捷的\x1b[31m狐\n狸跳过\x1b[0m了一只\n懒狗。",
			12,
		},
		// Handle mixed wide and short characters
		{
			"敏捷 A quick 的狐狸 fox 跳过 jumps over a lazy 了一只懒狗 dog。",
			"敏捷 A quick\n的狐狸 fox\n跳过 jumps\nover a lazy\n了一只懒狗\ndog。",
			12,
		},
		// Handle mixed wide and short characters with color
		{
			"敏捷 A \x1b31mquick 的狐狸 fox 跳\x1b0m过 jumps over a lazy 了一只懒狗 dog。",
			"敏捷 A \x1b31mquick\n的狐狸 fox\n跳\x1b0m过 jumps\nover a lazy\n了一只懒狗\ndog。",
			12,
		},
	}

	for i, tc := range cases {
		actual, lines := Wrap(tc.Input, tc.Lim)
		if actual != tc.Output {
			t.Fatalf("Case %d Input:\n\n`%s`\n\nExpected Output:\n\n`%s`\n\nActual Output:\n\n`%s`",
				i, tc.Input, tc.Output, actual)
		}

		expected := len(strings.Split(tc.Output, "\n"))
		if expected != lines {
			t.Fatalf("Case %d Nb lines mismatch\nExpected:%d\nActual:%d",
				i, expected, lines)
		}
	}
}

func TestWrapLeftPadded(t *testing.T) {
	cases := []struct {
		input, output string
		lim, pad      int
	}{
		{
			"The Lorem ipsum text is typically composed of pseudo-Latin words. It is commonly used as placeholder text to examine or demonstrate the visual effects of various graphic design.",
			`    The Lorem ipsum text is typically composed of
    pseudo-Latin words. It is commonly used as placeholder
    text to examine or demonstrate the visual effects of
    various graphic design.`,
			59, 4,
		},
		// Handle Chinese
		{
			"婞一枳郲逴靲屮蜧曀殳，掫乇峔掮傎溒兀緉冘仜。郼牪艽螗媷錵朸一詅掜豗怙刉笀丌，楀棶乇矹迡搦囷圣亍昄漚粁仈祂。覂一洳袶揙楱亍滻瘯毌，掗屮柅軡菵腩乜榵毌夯。勼哻怌婇怤灟葠雺奷朾恦扰衪岨坋誁乇芚誙腞。冇笉妺悆浂鱦賌廌灱灱觓坋佫呬耴跣兀枔蓔輈。嵅咍犴膰痭瘰机一靬涽捊矷尒玶乇，煚塈丌岰陊鉖怞戉兀甿跾觓夬侄。棩岧汌橩僁螗玎一逭舴圂衪扐衲兀，嵲媕亍衩衿溽昃夯丌侄蒰扂丱呤。毰侘妅錣廇螉仴一暀淖蚗佶庂咺丌，輀鈁乇彽洢溦洰氶乇构碨洐巿阹。",
			`    婞一枳郲逴靲屮蜧曀殳，掫乇峔掮傎溒兀緉冘仜。郼牪艽螗媷
    錵朸一詅掜豗怙刉笀丌，楀棶乇矹迡搦囷圣亍昄漚粁仈祂。覂
    一洳袶揙楱亍滻瘯毌，掗屮柅軡菵腩乜榵毌夯。勼哻怌婇怤灟
    葠雺奷朾恦扰衪岨坋誁乇芚誙腞。冇笉妺悆浂鱦賌廌灱灱觓坋
    佫呬耴跣兀枔蓔輈。嵅咍犴膰痭瘰机一靬涽捊矷尒玶乇，煚塈
    丌岰陊鉖怞戉兀甿跾觓夬侄。棩岧汌橩僁螗玎一逭舴圂衪扐衲
    兀，嵲媕亍衩衿溽昃夯丌侄蒰扂丱呤。毰侘妅錣廇螉仴一暀淖
    蚗佶庂咺丌，輀鈁乇彽洢溦洰氶乇构碨洐巿阹。`,
			59, 4,
		},
		// Handle long unbreakable words in a full stentence
		{
			"OT: there are alternatives to maintainer-/user-set priority, e.g. \"[user pain](http://www.lostgarden.com/2008/05/improving-bug-triage-with-user-pain.html)\".",
			`    OT: there are alternatives to maintainer-/user-set
    priority, e.g. "[user pain](http://www.lostgarden.com/
    2008/05/improving-bug-triage-with-user-pain.html)".`,
			58, 4,
		},
	}

	for i, tc := range cases {
		actual, lines := WrapLeftPadded(tc.input, tc.lim, tc.pad)
		if actual != tc.output {
			t.Fatalf("Case %d Input:\n\n`%s`\n\nExpected Output:\n`\n%s`\n\nActual Output:\n`\n%s\n%s`",
				i, tc.input, tc.output,
				"|"+strings.Repeat("-", tc.lim-2)+"|",
				actual)
		}

		expected := len(strings.Split(tc.output, "\n"))
		if expected != lines {
			t.Fatalf("Case %d Nb lines mismatch\nExpected:%d\nActual:%d",
				i, expected, lines)
		}
	}
}

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
	}

	for i, tc := range cases {
		l := wordLen(tc.Input)
		if l != tc.Length {
			t.Fatalf("Case %d Input:\n\n`%s`\n\nExpected Output:\n\n`%d`\n\nActual Output:\n\n`%d`",
				i, tc.Input, tc.Length, l)
		}
	}
}

func TestSplitWord(t *testing.T) {
	cases := []struct {
		Input            string
		Length           int
		Result, Leftover string
	}{
		// A simple word passes through.
		{
			"foo",
			4,
			"foo", "",
		},
		// Cut at the right place
		{
			"foobarHoy",
			4,
			"foob", "arHoy",
		},
		// A simple word passes through with colors
		{
			"\x1b[31mbar\x1b[0m",
			4,
			"\x1b[31mbar\x1b[0m", "",
		},
		// Cut at the right place with colors
		{
			"\x1b[31mfoobarHoy\x1b[0m",
			4,
			"\x1b[31mfoob", "arHoy\x1b[0m",
		},
		// Handle prefix and suffix properly
		{
			"foo\x1b[31mfoobarHoy\x1b[0mbaaar",
			4,
			"foo\x1b[31mf", "oobarHoy\x1b[0mbaaar",
		},
		// Cut properly with length = 0
		{
			"foo",
			0,
			"", "foo",
		},
		// Handle chinese
		{
			"快檢什麼望對",
			4,
			"快檢", "什麼望對",
		},
		{
			"快檢什麼望對",
			5,
			"快檢", "什麼望對",
		},
		// Handle chinese with colors
		{
			"快\x1b[31m檢什麼\x1b[0m望對",
			4,
			"快\x1b[31m檢", "什麼\x1b[0m望對",
		},
	}

	for i, tc := range cases {
		result, leftover := splitWord(tc.Input, tc.Length)
		if result != tc.Result || leftover != tc.Leftover {
			t.Fatalf("Case %d Input:\n\n`%s`\n\nExpected Output:\n\n`%s` - `%s`\n\nActual Output:\n\n`%s` - `%s`",
				i, tc.Input, tc.Result, tc.Leftover, result, leftover)
		}
	}
}

func TestExtractApplyTermEscapes(t *testing.T) {
	cases := []struct {
		Name        string
		Input       string
		Output      string
		TermEscapes []escapeItem
	}{

		{
			"A plain ascii line with escapes",
			"This \x1b[31mis an\x1b[0m example.",
			"This is an example.",
			[]escapeItem{{"\x1b[31m", 5}, {"\x1b[0m", 10}},
		},

		{
			"Escape at the end",
			"This \x1b[31mis an example.\x1b[0m",
			"This is an example.",
			[]escapeItem{{"\x1b[31m", 5}, {"\x1b[0m", 19}},
		},

		{
			"A plain wide line with escapes",
			"一只敏捷\x1b[31m的狐狸\x1b[0m跳过了一只懒狗。",
			"一只敏捷的狐狸跳过了一只懒狗。",
			[]escapeItem{{"\x1b[31m", 4}, {"\x1b[0m", 7}},
		},

		{
			"A normal-wide mixed line with escapes",
			"一只 A Quick 敏捷\x1b[31m的狐 Fox 狸\x1b[0m跳过了Dog一只懒狗。",
			"一只 A Quick 敏捷的狐 Fox 狸跳过了Dog一只懒狗。",
			[]escapeItem{{"\x1b[31m", 13}, {"\x1b[0m", 21}},
		},

		{
			"Multiple escapes at the same place",
			"\x1b[1m\x1b[31mThis \x1b[1m\x1b[31mis an\x1b[0m example.\x1b[1m\x1b[31m",
			"This is an example.",
			[]escapeItem{
				{"\x1b[1m", 0}, {"\x1b[31m", 0},
				{"\x1b[1m", 5}, {"\x1b[31m", 5},
				{"\x1b[0m", 10},
				{"\x1b[1m", 19}, {"\x1b[31m", 19}},
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			cleaned, escapes := extractTermEscapes(tc.Input)

			assert.Equal(t, tc.Output, cleaned)
			assert.Equal(t, tc.TermEscapes, escapes)

			augmented := applyTermEscapes(cleaned, escapes)

			assert.Equal(t, tc.Input, augmented)
		})
	}
}

func TestSegmentLines(t *testing.T) {
	cases := []struct {
		Input  string
		Output []string
	}{
		// A plain ascii line with escapes.
		{
			"This is an example.",
			[]string{"This", " ", "is", " ", "an", " ", "example."},
		},
		// A plain wide line with escapes.
		{
			"一只敏捷的狐狸跳过了一只懒狗。",
			[]string{"一", "只", "敏", "捷", "的", "狐", "狸", "跳", "过",
				"了", "一", "只", "懒", "狗", "。"},
		},
		// A complex stentence.
		{
			"This is a 'complex' example, where   一只 and English 混合了。",
			[]string{"This", " ", "is", " ", "a", " ", "'complex'", " ", "example,",
				" ", "where", "   ", "一", "只", " ", "and", " ", "English", " ", "混",
				"合", "了", "。"},
		},
	}

	for i, tc := range cases {
		chunks := segmentLine(tc.Input)
		if !reflect.DeepEqual(chunks, tc.Output) {
			t.Fatalf("Case %d Input:\n\n`%s`\n\nExpected Output:\n\n`[%s]`\n\nActual Output:\n\n`[%s]`\n\n",
				i, tc.Input, strings.Join(tc.Output, ", "), strings.Join(chunks, ", "))
		}
	}
}
