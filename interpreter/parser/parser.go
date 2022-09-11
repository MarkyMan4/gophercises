package parser

import (
	"github.com/MarkyMan4/simple-interpreter/ast"
	"github.com/MarkyMan4/simple-interpreter/lexer"
	"github.com/MarkyMan4/simple-interpreter/token"
)

type Parser struct {
	Lex       *lexer.Lexer
	prevToken token.Token
	curToken  token.Token
	peekToken token.Token
}

func NewParser(l *lexer.Lexer) *Parser {
	p := &Parser{Lex: l}

	// load two tokens
	// prevToken will be nil, curToken and peek token will have the first two tokens
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.prevToken = p.curToken
	p.curToken = p.peekToken
	p.peekToken = p.Lex.NextToken()
}

func (p *Parser) Parse() []*ast.Node {
	return nil
}
