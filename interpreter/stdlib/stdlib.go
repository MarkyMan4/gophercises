package stdlib

import (
	"fmt"

	"github.com/MarkyMan4/simple-interpreter/object"
)

type BuiltIn func(args ...object.Object) object.Object

var BuiltInFuns = map[string]BuiltIn{
	"print": PrintFun,
}

func PrintFun(args ...object.Object) object.Object {
	// print each argument separated by space and ending with a newline
	for i := range args {
		fmt.Print(args[i].ToString())

		if i == len(args)-1 {
			fmt.Println()
		} else {
			fmt.Print(" ")
		}
	}

	return &object.NullObject{}
}
