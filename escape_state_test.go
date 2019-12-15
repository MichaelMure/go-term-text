package text

import "testing"

func TestEscapeState(t *testing.T) {
	cases := []struct {
		input  string
		output string
	}{
		{
			"Format: ![Alt Text](\x1b[34murl\x1b[0m)",
			"\x1b[0m",
		},
		{
			"\x1b[1;2;3;4;5;7;8;9;33;43m",
			"\x1b[1;2;3;4;5;7;8;9;33;43m",
		},
		{
			"baaar\x1b[48;5;118mfoobar\x1b[38;5;100m",
			"\x1b[38;5;100;48;5;118m",
		},
		{
			"baaar\x1b[48;2;118;131;193mfoobar\x1b[38;2;255;255;250mfooooo",
			"\x1b[38;2;255;255;250;48;2;118;131;193m",
		},
	}

	for i, tc := range cases {
		es := &EscapeState{}

		es.Witness(tc.input)

		result := es.String()
		if result != tc.output {
			t.Fatalf("Case %d Input:\n\n`%s`\n\nExpected Output:\n\n`%s`\n\nActual Output:\n\n`%s`",
				i, tc.input, tc.output, result)
		}
	}
}
