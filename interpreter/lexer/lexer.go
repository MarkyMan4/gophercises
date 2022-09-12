package lexer

import (
	"unicode"

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
	if l.readPos >= len(l.chars) {
		l.curChar = rune(0)
	} else {
		l.curChar = l.chars[l.readPos]
	}

	l.curPos = l.readPos
	l.readPos++
}

// peek ahead one character without increasing curPos or readPos
func (l *Lexer) peek() rune {
	var peekedChar rune

	if l.readPos >= len(l.chars) {
		peekedChar = rune(0)
	} else {
		peekedChar = l.chars[l.readPos]
	}

	return peekedChar
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
	case '=':
		tok = token.Token{Type: token.ASSIGN, Literal: string(l.curChar)}
	case '(':
		tok = token.Token{Type: token.LPAREN, Literal: string(l.curChar)}
	case ')':
		tok = token.Token{Type: token.RPAREN, Literal: string(l.curChar)}
	case '{':
		tok = token.Token{Type: token.RBRACE, Literal: string(l.curChar)}
	case '}':
		tok = token.Token{Type: token.LBRACE, Literal: string(l.curChar)}
	case '[':
		tok = token.Token{Type: token.RBRACK, Literal: string(l.curChar)}
	case ']':
		tok = token.Token{Type: token.LBRACK, Literal: string(l.curChar)}
	case ';':
		tok = token.Token{Type: token.SEMI, Literal: string(l.curChar)}
	case rune(0):
		tok = token.Token{Type: token.EOF, Literal: ""}
	default:
		// read a number or an identifier
		if unicode.IsDigit(l.curChar) {
			tok = l.readIntOrFloat()
		} else {
			tok = l.readIdent()
		}
	}

	l.nextChar()
	return tok
}

func (l *Lexer) skipWhitespace() {
	for l.curChar == ' ' || l.curChar == '\n' || l.curChar == '\t' || l.curChar == '\r' {
		l.nextChar()
	}
}

func (l *Lexer) readIntOrFloat() token.Token {
	literal := l.readNumber()
	var tokType string

	// if a decimal comes after the number, read the decimal and then all the numbers after it - it's a float
	if l.peek() == '.' {
		l.nextChar()
		literal += string(l.curChar)
		l.nextChar() // read past decimal before reading digits after decimal
		literal += l.readNumber()
		tokType = token.FLOAT
	} else {
		tokType = token.INT
	}

	return token.Token{Type: tokType, Literal: literal}
}

func (l *Lexer) readNumber() string {
	numTok := string(l.curChar)

	for unicode.IsDigit(l.peek()) {
		l.nextChar()
		numTok += string(l.curChar)
	}

	return numTok
}

func (l *Lexer) readIdent() token.Token {
	literal := string(l.curChar)

	for unicode.IsLetter(l.peek()) || unicode.IsDigit(l.peek()) {
		l.nextChar()
		literal += string(l.curChar)
	}

	return token.Token{Type: token.GetIdentOrKeyword(literal), Literal: literal}
}
