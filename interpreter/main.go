package main

import (
	"fmt"

	"github.com/MarkyMan4/simple-interpreter/ast"
	"github.com/MarkyMan4/simple-interpreter/object"
)

func eval(node ast.Node, env *object.Environment) object.Object {
	switch node := node.(type) {
	case *ast.IntegerLiteral:
		return &object.IntegerObject{Value: node.Value}
	case *ast.FloatLiteral:
		return &object.FloatObject{Value: node.Value}
	case *ast.StringLiteral:
		return &object.StringObject{Value: node.Value}
	case *ast.IdentifierExpression:
		return env.Get(node.Value)
	case *ast.InfixExpression:
		left := eval(node.Left, env)
		right := eval(node.Right, env)
		return evalInfixExpression(node.Op, left, right)
	case *ast.LetStatement:
		val := eval(node.Value, env)
		env.Set(node.Identifier, val)
	}

	return nil
}

func evalInfixExpression(op ast.Operator, left object.Object, right object.Object) object.Object {
	if left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ {
		return evalIntegerInfixExpression(op, left, right)
	} else if left.Type() == object.INTEGER_OBJ && right.Type() == object.FLOAT_OBJ {
		return evalIntegerFloatInfixExpression(op, left, right)
	} else if left.Type() == object.FLOAT_OBJ && right.Type() == object.INTEGER_OBJ {
		return evalFloatIntegerInfixExpression(op, left, right)
	} else if left.Type() == object.FLOAT_OBJ && right.Type() == object.FLOAT_OBJ {
		return evalFloatInfixExpression(op, left, right)
	}

	return nil
}

func evalIntegerInfixExpression(op ast.Operator, left object.Object, right object.Object) object.Object {
	leftVal := left.(*object.IntegerObject).Value
	rightVal := right.(*object.IntegerObject).Value

	switch op {
	case ast.Plus:
		return &object.IntegerObject{Value: leftVal + rightVal}
	case ast.Minus:
		return &object.IntegerObject{Value: leftVal - rightVal}
	case ast.Mult:
		return &object.IntegerObject{Value: leftVal * rightVal}
	case ast.Divide:
		// for dividing integers, convert them to floats so we get a float in return
		return &object.FloatObject{Value: float64(leftVal) / float64(rightVal)}
	default:
		return &object.ErrorObject{Message: fmt.Sprintf("unsupported operator '%c' for types %s, %s", op, left.Type(), right.Type())}
	}
}

func evalFloatIntegerInfixExpression(op ast.Operator, left object.Object, right object.Object) object.Object {
	leftVal := left.(*object.FloatObject).Value
	rightVal := float64(right.(*object.IntegerObject).Value)

	switch op {
	case ast.Plus:
		return &object.FloatObject{Value: leftVal + rightVal}
	case ast.Minus:
		return &object.FloatObject{Value: leftVal - rightVal}
	case ast.Mult:
		return &object.FloatObject{Value: leftVal * rightVal}
	case ast.Divide:
		return &object.FloatObject{Value: leftVal / rightVal}
	default:
		return &object.ErrorObject{Message: fmt.Sprintf("unsupported operator '%c' for types %s, %s", op, left.Type(), right.Type())}
	}
}

func evalIntegerFloatInfixExpression(op ast.Operator, left object.Object, right object.Object) object.Object {
	leftVal := float64(left.(*object.IntegerObject).Value)
	rightVal := right.(*object.FloatObject).Value

	switch op {
	case ast.Plus:
		return &object.FloatObject{Value: leftVal + rightVal}
	case ast.Minus:
		return &object.FloatObject{Value: leftVal - rightVal}
	case ast.Mult:
		return &object.FloatObject{Value: leftVal * rightVal}
	case ast.Divide:
		return &object.FloatObject{Value: leftVal / rightVal}
	default:
		return &object.ErrorObject{Message: fmt.Sprintf("unsupported operator '%c' for types %s, %s", op, left.Type(), right.Type())}
	}
}

func evalFloatInfixExpression(op ast.Operator, left object.Object, right object.Object) object.Object {
	leftVal := left.(*object.FloatObject).Value
	rightVal := right.(*object.FloatObject).Value

	switch op {
	case ast.Plus:
		return &object.FloatObject{Value: leftVal + rightVal}
	case ast.Minus:
		return &object.FloatObject{Value: leftVal - rightVal}
	case ast.Mult:
		return &object.FloatObject{Value: leftVal * rightVal}
	case ast.Divide:
		return &object.FloatObject{Value: leftVal / rightVal}
	default:
		return &object.ErrorObject{Message: fmt.Sprintf("unsupported operator '%c' for types %s, %s", op, left.Type(), right.Type())}
	}
}

func main() {
	env := object.NewEnvironment()

	// (1 + (2 / 3.456)) * 2
	// res = 3.157407
	expr1 := &ast.InfixExpression{
		Left: &ast.InfixExpression{
			Left: &ast.IntegerLiteral{Value: 1},
			Op:   ast.Plus,
			Right: &ast.InfixExpression{
				Left:  &ast.IntegerLiteral{Value: 2},
				Op:    ast.Divide,
				Right: &ast.FloatLiteral{Value: 3.456},
			},
		},
		Op:    ast.Mult,
		Right: &ast.IntegerLiteral{Value: 2},
	}

	res := eval(expr1, env)
	fmt.Println(res.ToString())

	// invalid operator $ for float and integer
	expr2 := &ast.InfixExpression{
		Left: &ast.InfixExpression{
			Left: &ast.IntegerLiteral{Value: 1},
			Op:   ast.Plus,
			Right: &ast.InfixExpression{
				Left:  &ast.IntegerLiteral{Value: 2},
				Op:    ast.Divide,
				Right: &ast.FloatLiteral{Value: 3.456},
			},
		},
		Op:    '$',
		Right: &ast.IntegerLiteral{Value: 2},
	}

	res = eval(expr2, env)
	fmt.Println(res.ToString())
}
