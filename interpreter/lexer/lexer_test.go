package lexer

import (
	"fmt"
	"testing"

	"github.com/MarkyMan4/simple-interpreter/token"
)

func TestVar(t *testing.T) {
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

func TestWhile(t *testing.T) {
	input := "var x = 1; while(x < 5) {x += 1;}"
	// input := "var x = true; x = 4;"
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
