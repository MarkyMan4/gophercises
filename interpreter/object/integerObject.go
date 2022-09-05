package object

import "fmt"

type IntegerObject struct {
	Value int
}

func (i *IntegerObject) Type() string {
	return INTEGER_OBJ
}

func (i *IntegerObject) ToString() string {
	return fmt.Sprint(i.Value)
}
