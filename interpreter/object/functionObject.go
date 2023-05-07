package object

import (
	"github.com/MarkyMan4/simple-interpreter/ast"
)

type FunctionObject struct {
	Args       []string
	Statements []ast.Statement
}

func (f *FunctionObject) Type() string {
	return FUNCTION_OBJ
}

func (f *FunctionObject) ToString() string {
	return "" // TODO: implement this later
}
