package ast

import "fmt"

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
	Value int64
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

func (b *StringLiteral) expressionNode() {}

type BooleanLiteral struct {
	Value bool
}

func (b *BooleanLiteral) ToString() string {
	return fmt.Sprint(b.Value)
}

func (b *BooleanLiteral) expressionNode() {}

type InfixExpression struct {
	Left  Expression
	Op    string
	Right Expression
}

func (i *InfixExpression) ToString() string {
	// fix this later, just returning empty string for now
	return i.Left.ToString() + i.Op + i.Right.ToString()
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
type VarStatement struct {
	Identifier string
	Value      Expression
}

func (l *VarStatement) ToString() string {
	return fmt.Sprintf("var %s = %s;", l.Identifier, l.Value.ToString())
}

func (l *VarStatement) statementNode() {}

// program is a list of statements
type Program struct {
	Statements []Statement
}

func (p *Program) ToString() string {
	return ""
}

// while loop
type WhileStatement struct {
	Condition  Expression
	Statements []Statement
}

func (ws *WhileStatement) ToString() string {
	whileAsStr := fmt.Sprintf("while(%s) {", ws.Condition.ToString())

	for i := range ws.Statements {
		whileAsStr += ws.Statements[i].ToString() + " "
	}

	whileAsStr += "}"

	return whileAsStr
}

func (ws *WhileStatement) statementNode() {}
