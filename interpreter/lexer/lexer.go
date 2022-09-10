package lexer

// import (
// 	"github.com/MarkyMan4/simple-interpreter/token"
// )

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
