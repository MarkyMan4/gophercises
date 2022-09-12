package parser

import (
	"fmt"
	"strconv"

	"github.com/MarkyMan4/simple-interpreter/ast"
	"github.com/MarkyMan4/simple-interpreter/lexer"
	"github.com/MarkyMan4/simple-interpreter/token"
)

type (
	prefixParser func() ast.Expression
	infixParser  func(ast.Expression) ast.Expression
)

type Parser struct {
	Lex           *lexer.Lexer
	prevToken     token.Token
	curToken      token.Token
	peekToken     token.Token
	Errors        []string
	prefixParsers map[string]prefixParser
	infixParsers  map[string]infixParser
}

func NewParser(l *lexer.Lexer) *Parser {
	p := &Parser{Lex: l}

	// load two tokens
	// prevToken will be nil, curToken and peek token will have the first two tokens
	p.nextToken()
	p.nextToken()

	// prefix parsers (e.g. an int is a prefix in an expression)
	p.prefixParsers = map[string]prefixParser{
		token.INT:   p.parseIntegerLiteral,
		token.FLOAT: p.parseFloatLiteral,
		token.IDENT: p.parseIdent,
	}

	// infix parsers (e.g. +, -, *, /)
	p.infixParsers = map[string]infixParser{
		token.PLUS:   p.parseInfixExpression,
		token.MINUS:  p.parseInfixExpression,
		token.MULT:   p.parseInfixExpression,
		token.DIVIDE: p.parseInfixExpression,
	}

	return p
}

func (p *Parser) nextToken() {
	p.prevToken = p.curToken
	p.curToken = p.peekToken
	p.peekToken = p.Lex.NextToken()
}

func (p *Parser) Parse() *ast.Program {
	prog := &ast.Program{Statements: []ast.Statement{}}

	for p.curToken.Type != token.EOF {
		prog.Statements = append(prog.Statements, p.parseStmt())
		p.nextToken()
	}

	return prog
}

func (p *Parser) parseStmt() ast.Statement {
	switch p.curToken.Type {
	case token.VAR:
		return p.parseVarStmt()
	default:
		return nil
	}
}

func (p *Parser) parseVarStmt() ast.Statement {
	if !p.expectNextToken(token.IDENT) {
		return nil
	}

	p.nextToken()
	stmt := &ast.VarStatement{Identifier: p.curToken.Literal}

	if !p.expectNextToken(token.ASSIGN) {
		return nil
	}

	p.nextToken()
	p.nextToken()
	stmt.Value = p.parseExpression()

	return stmt
}

func (p *Parser) expectNextToken(tokType string) bool {
	if p.peekToken.Type == tokType {
		return true
	}

	errMsg := fmt.Sprintf("Expected token %s to be %s, but got %s", p.peekToken, tokType, p.peekToken.Type)
	p.Errors = append(p.Errors, errMsg)

	return false
}

func (p *Parser) parseExpression() ast.Expression {
	prefix := p.prefixParsers[p.curToken.Type]
	if prefix == nil {
		// TODO: do some kind of error here
		return nil
	}

	left := prefix()
	for p.peekToken.Type != token.SEMI {
		infix := p.infixParsers[p.peekToken.Type]
		if infix == nil {
			return left
		}

		p.nextToken()
		left = infix(left)
	}

	return left
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	intLit := &ast.IntegerLiteral{}
	val, err := strconv.ParseInt(p.curToken.Literal, 10, 64)

	if err != nil {
		errMsg := fmt.Sprintf("could not parse %s as type integer", p.curToken.Literal)
		p.Errors = append(p.Errors, errMsg)
		return nil
	}

	intLit.Value = val

	return intLit
}

func (p *Parser) parseFloatLiteral() ast.Expression {
	floatLit := &ast.FloatLiteral{}
	val, err := strconv.ParseFloat(p.curToken.Literal, 64)

	if err != nil {
		errMsg := fmt.Sprintf("could not parse %s as type float", p.curToken.Literal)
		p.Errors = append(p.Errors, errMsg)
		return nil
	}

	floatLit.Value = val

	return floatLit
}

func (p *Parser) parseIdent() ast.Expression {
	return &ast.IdentifierExpression{Value: p.curToken.Literal}
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expr := &ast.InfixExpression{
		Op:   p.curToken.Literal,
		Left: left,
	}

	p.nextToken()
	expr.Right = p.parseExpression()

	return expr
}
