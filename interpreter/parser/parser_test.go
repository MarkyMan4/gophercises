package parser

import (
	"fmt"
	"testing"

	"github.com/MarkyMan4/simple-interpreter/lexer"
)

func TestParse(t *testing.T) {
	l := lexer.NewLexer("var y = true; while(1 + 2) {var x = 3;}")
	p := NewParser(l)
	prog := p.Parse()

	stmt := prog.Statements[0]

	fmt.Println(stmt.ToString())

	// tok := l.NextToken()

	// for tok.Type != token.EOF {
	// 	fmt.Println(tok)
	// 	tok = l.NextToken()
	// }
}
