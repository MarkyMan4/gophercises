package lexer

import (
	"github.com/MarkyMan4/simple-interpreter/token"
)

type Lexer struct {
	Tokens  []token.Token
	curPos  int
	nextPos int
}
