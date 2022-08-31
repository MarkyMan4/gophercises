package main

import (
	"fmt"
)

type Operator rune
type Number int

const (
	Plus   Operator = '+'
	Minus           = '-'
	Mult            = '*'
	Divide          = '/'
)

type Expr interface {
	eval() int
}

type ComplexExpr struct {
	Op    Operator
	Expr1 Expr
	Expr2 Expr
}

type SimpleExpr struct {
	Num Number
}

func (e *ComplexExpr) eval() int {
	result := 0

	lhsRes := e.Expr1.eval()
	rhsRes := e.Expr2.eval()

	switch e.Op {
	case '+':
		result = lhsRes + rhsRes
	case '-':
		result = lhsRes - rhsRes
	case '*':
		result = lhsRes * rhsRes
	case '/':
		result = lhsRes / rhsRes
	}

	return result
}

func (e *SimpleExpr) eval() int {
	return int(e.Num)
}

func main() {
	// 3 + ((4 * 5) - 2)
	// res = 21
	expr := ComplexExpr{
		Plus,
		&SimpleExpr{3},
		&ComplexExpr{
			Minus,
			&ComplexExpr{
				Mult,
				&SimpleExpr{4},
				&SimpleExpr{5},
			},
			&SimpleExpr{2},
		},
	}

	res := expr.eval()
	fmt.Println(res)
}
