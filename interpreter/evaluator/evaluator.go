package evaluator

import (
	"fmt"
	"os"

	"github.com/MarkyMan4/simple-interpreter/ast"
	"github.com/MarkyMan4/simple-interpreter/object"
)

func Eval(node ast.Node, env *object.Environment) object.Object {
	switch node := node.(type) {
	case *ast.IntegerLiteral:
		return &object.IntegerObject{Value: node.Value}
	case *ast.FloatLiteral:
		return &object.FloatObject{Value: node.Value}
	case *ast.StringLiteral:
		return &object.StringObject{Value: node.Value}
	case *ast.BooleanLiteral:
		return &object.BooleanObject{Value: node.Value}
	case *ast.IdentifierExpression:
		return env.Get(node.Value)
	case *ast.InfixExpression:
		left := Eval(node.Left, env)
		right := Eval(node.Right, env)
		return evalInfixExpression(node.Op, left, right)
	case *ast.VarStatement:
		val := Eval(node.Value, env)
		env.Set(node.Identifier, val)
	case *ast.AssignStatement:
		if env.Get(node.Identifier) == nil {
			fmt.Printf("variable %s has not been declared\n", node.Identifier)
			os.Exit(1)
		}

		left := env.Get(node.Identifier)
		right := Eval(node.Value, env)
		val := evalAssignStatement(node.AssignOp, left, right)
		env.Set(node.Identifier, val)
	case *ast.IfStatement:
		condResult := Eval(node.Condition, env)
		if condResult.Type() != object.BOOLEAN_OBJ {
			fmt.Println("condition must return a boolean")
			os.Exit(1)
		}

		// if the condition is still true, run all statements and evaluate the loop again
		if condResult.(*object.BooleanObject).Value {
			for i := range node.Statements {
				Eval(node.Statements[i], env)
			}
		}
	case *ast.WhileStatement:
		condResult := Eval(node.Condition, env)
		if condResult.Type() != object.BOOLEAN_OBJ {
			fmt.Println("condition must return a boolean")
			os.Exit(1)
		}

		// if the condition is still true, run all statements and evaluate the loop again
		if condResult.(*object.BooleanObject).Value {
			for i := range node.Statements {
				Eval(node.Statements[i], env)
			}

			Eval(node, env)
		}
		// case *ast.FunctionDef:

	}

	return nil
}

func evalAssignStatement(assignOp string, left object.Object, right object.Object) object.Object {
	switch assignOp {
	case "+=":
		return evalInfixExpression("+", left, right)
	case "-=":
		return evalInfixExpression("-", left, right)
	case "*=":
		return evalInfixExpression("*", left, right)
	case "/=":
		return evalInfixExpression("/", left, right)
	case "=":
		// default is a normal assignment, so just return the right hand side
		return right
	default:
		fmt.Printf("unknown operator %s\n", assignOp)
		os.Exit(1)
	}

	return nil
}

func evalInfixExpression(op string, left object.Object, right object.Object) object.Object {
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

func evalIntegerInfixExpression(op string, left object.Object, right object.Object) object.Object {
	leftVal := left.(*object.IntegerObject).Value
	rightVal := right.(*object.IntegerObject).Value

	switch op {
	case "+":
		return &object.IntegerObject{Value: leftVal + rightVal}
	case "-":
		return &object.IntegerObject{Value: leftVal - rightVal}
	case "*":
		return &object.IntegerObject{Value: leftVal * rightVal}
	case "/":
		// for dividing integers, convert them to floats so we get a float in return
		return &object.FloatObject{Value: float64(leftVal) / float64(rightVal)}
	case "<":
		return &object.BooleanObject{Value: leftVal < rightVal}
	case "<=":
		return &object.BooleanObject{Value: leftVal <= rightVal}
	case "==":
		return &object.BooleanObject{Value: leftVal == rightVal}
	case ">":
		return &object.BooleanObject{Value: leftVal > rightVal}
	case ">=":
		return &object.BooleanObject{Value: leftVal >= rightVal}
	default:
		return &object.ErrorObject{Message: fmt.Sprintf("unsupported operator '%s' for types %s, %s", op, left.Type(), right.Type())}
	}
}

func evalFloatIntegerInfixExpression(op string, left object.Object, right object.Object) object.Object {
	leftVal := left.(*object.FloatObject).Value
	rightVal := float64(right.(*object.IntegerObject).Value)

	switch op {
	case "+":
		return &object.FloatObject{Value: leftVal + rightVal}
	case "-":
		return &object.FloatObject{Value: leftVal - rightVal}
	case "*":
		return &object.FloatObject{Value: leftVal * rightVal}
	case "/":
		return &object.FloatObject{Value: leftVal / rightVal}
	case "<":
		return &object.BooleanObject{Value: leftVal < rightVal}
	case "<=":
		return &object.BooleanObject{Value: leftVal <= rightVal}
	case "==":
		return &object.BooleanObject{Value: leftVal == rightVal}
	case ">":
		return &object.BooleanObject{Value: leftVal > rightVal}
	case ">=":
		return &object.BooleanObject{Value: leftVal >= rightVal}
	default:
		return &object.ErrorObject{Message: fmt.Sprintf("unsupported operator '%s' for types %s, %s", op, left.Type(), right.Type())}
	}
}

func evalIntegerFloatInfixExpression(op string, left object.Object, right object.Object) object.Object {
	leftVal := float64(left.(*object.IntegerObject).Value)
	rightVal := right.(*object.FloatObject).Value

	switch op {
	case "+":
		return &object.FloatObject{Value: leftVal + rightVal}
	case "-":
		return &object.FloatObject{Value: leftVal - rightVal}
	case "*":
		return &object.FloatObject{Value: leftVal * rightVal}
	case "/":
		return &object.FloatObject{Value: leftVal / rightVal}
	case "<":
		return &object.BooleanObject{Value: leftVal < rightVal}
	case "<=":
		return &object.BooleanObject{Value: leftVal <= rightVal}
	case "==":
		return &object.BooleanObject{Value: leftVal == rightVal}
	case ">":
		return &object.BooleanObject{Value: leftVal > rightVal}
	case ">=":
		return &object.BooleanObject{Value: leftVal >= rightVal}
	default:
		return &object.ErrorObject{Message: fmt.Sprintf("unsupported operator '%s' for types %s, %s", op, left.Type(), right.Type())}
	}
}

func evalFloatInfixExpression(op string, left object.Object, right object.Object) object.Object {
	leftVal := left.(*object.FloatObject).Value
	rightVal := right.(*object.FloatObject).Value

	switch op {
	case "+":
		return &object.FloatObject{Value: leftVal + rightVal}
	case "-":
		return &object.FloatObject{Value: leftVal - rightVal}
	case "*":
		return &object.FloatObject{Value: leftVal * rightVal}
	case "/":
		return &object.FloatObject{Value: leftVal / rightVal}
	case "<":
		return &object.BooleanObject{Value: leftVal < rightVal}
	case "<=":
		return &object.BooleanObject{Value: leftVal <= rightVal}
	case "==":
		return &object.BooleanObject{Value: leftVal == rightVal}
	case ">":
		return &object.BooleanObject{Value: leftVal > rightVal}
	case ">=":
		return &object.BooleanObject{Value: leftVal >= rightVal}
	default:
		return &object.ErrorObject{Message: fmt.Sprintf("unsupported operator '%s' for types %s, %s", op, left.Type(), right.Type())}
	}
}
