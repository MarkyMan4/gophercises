package lexer

import (
	"fmt"
	"testing"
)

func TestPrint(t *testing.T) {
	// converting string to slice of runes
	s := "this is a test"
	chars := []rune(s)

	for i := range chars {
		fmt.Println(string(chars[i]))
	}
}
