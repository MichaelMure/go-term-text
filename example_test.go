package text

import (
	"fmt"
	"strings"
)

func ExampleWrapWithPadIndent() {
	input := "The \x1b[1mLorem ipsum\x1b[0m text is typically composed of " +
		"pseudo-Latin words. It is commonly used as \x1b[3mplaceholder\x1b[0m" +
		" text to examine or demonstrate the \x1b[9mvisual effects\x1b[0m of " +
		"various graphic design. 一只 A Quick \x1b[31m敏捷的狐 Fox " +
		"狸跳过了\x1b[0mDog一只懒狗。"

	output, n := WrapWithPadIndent(input, 60,
		"\x1b[34m<-indent-> \x1b[0m", "\x1b[33m<-pad-> \x1b[0m")

	fmt.Println()
	fmt.Printf("output has %d lines\n\n", n)

	fmt.Println("|" + strings.Repeat("-", 58) + "|")
	fmt.Println(output)
	fmt.Println("|" + strings.Repeat("-", 58) + "|")
	fmt.Println()
}
