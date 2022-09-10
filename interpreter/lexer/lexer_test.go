package lexer

import (
	"fmt"
	"testing"

	"github.com/MarkyMan4/simple-interpreter/token"
)

func TestPrint(t *testing.T) {
	input := "var x = 5; var y = 10.123 + 90;"
	lex := NewLexer(input)

	tok := lex.NextToken()
	tokens := []token.Token{tok}

	for tok.Type != token.EOF {
		tok = lex.NextToken()
		tokens = append(tokens, tok)
	}

	fmt.Println(input + "\n")

	for i := range tokens {
		fmt.Println(tokens[i])
	}
}
