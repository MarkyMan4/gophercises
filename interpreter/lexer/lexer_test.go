package lexer

import (
	"fmt"
	"testing"
)

func TestPrint(t *testing.T) {
	var m = map[string]string{
		"var": "var",
	}

	tok, ok := m["var"]
	fmt.Println(ok)

	if ok {
		fmt.Println(tok)
	}
}
