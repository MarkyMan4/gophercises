package lexer

import (
	"github.com/MarkyMan4/simple-interpreter/token"
)

type Lexer struct {
	curPos  int
	readPos int
	curChar rune
	chars   []rune
}

func NewLexer(input string) *Lexer {
	l := &Lexer{chars: []rune(input)}
	l.nextChar()
	return l
}

func (l *Lexer) nextChar() {
	l.curChar = l.chars[l.readPos]
	l.curPos = l.readPos
	l.readPos++
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhitespace()

	switch l.curChar {
	case '+':
		tok = token.Token{Type: token.PLUS, Literal: string(l.curChar)}
	case '-':
		tok = token.Token{Type: token.MINUS, Literal: string(l.curChar)}
	case '*':
		tok = token.Token{Type: token.MULT, Literal: string(l.curChar)}
	case '/':
		tok = token.Token{Type: token.DIVIDE, Literal: string(l.curChar)}
	case '(':
		tok = token.Token{Type: token.LPAREN, Literal: string(l.curChar)}
	case ')':
		tok = token.Token{Type: token.RPAREN, Literal: string(l.curChar)}
	case ';':
		tok = token.Token{Type: token.SEMI, Literal: string(l.curChar)}
	default:

	}

	return tok
}

func (l *Lexer) skipWhitespace() {
	for l.curChar == ' ' || l.curChar == '\n' || l.curChar == '\t' || l.curChar == '\r' {
		l.nextChar()
	}
}
