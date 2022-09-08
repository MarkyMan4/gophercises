package ast

import "fmt"

type Operator rune

const (
	Plus   Operator = '+'
	Minus           = '-'
	Mult            = '*'
	Divide          = '/'
)

type Node interface {
	ToString() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// expressions
type IntegerLiteral struct {
	Value int
}

func (i *IntegerLiteral) ToString() string {
	return fmt.Sprint(i.Value)
}

func (i *IntegerLiteral) expressionNode() {}

type FloatLiteral struct {
	Value float64
}

func (i *FloatLiteral) ToString() string {
	return fmt.Sprint(i.Value)
}

func (i *FloatLiteral) expressionNode() {}

type StringLiteral struct {
	Value string
}

func (i *StringLiteral) ToString() string {
	return i.Value
}

func (i *StringLiteral) expressionNode() {}

type InfixExpression struct {
	Left  Expression
	Op    Operator
	Right Expression
}

func (i *InfixExpression) ToString() string {
	// fix this later, just returning empty string for now
	return ""
}

func (i *InfixExpression) expressionNode() {}

type IdentifierExpression struct {
	Value string // name of identifier
}

func (i *IdentifierExpression) ToString() string {
	return i.Value
}

func (i *IdentifierExpression) expressionNode() {}

// statements
type LetStatement struct {
	Identifier string
	Value      Expression
}

func (l *LetStatement) ToString() string {
	return fmt.Sprintf("let %s = %s", l.Identifier, l.Value.ToString())
}

func (l *LetStatement) statementNode() {}
