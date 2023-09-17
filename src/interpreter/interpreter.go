package src

import (
	"fmt"
	"reflect"
)

func getNode(ast interface{}, key string) interface{} {
	fmt.Println(ast.(map[string]interface{})[key])
	return ast.(map[string]interface{})[key]
}

func copyEnvironment(environment map[string]interface{}) map[string]interface{} {
	newEnvironment := make(map[string]interface{})
	for key, value := range environment {
		newEnvironment[key] = value
	}
	return newEnvironment
}

func Execute(ast interface{}, environment map[string]interface{}) interface{} {
	switch kind := getNode(ast, "kind"); kind {
	case PRINT:
		term := Execute(getNode(ast, "value"), environment)
		fmt.Println(term)
		return nil
	case STR:
		return getNode(ast, "value")
	case INT:
		return getNode(ast, "value")
	case BINARY:
		lhs := Execute(getNode(ast, "lhs"), environment)
		rhs := Execute(getNode(ast, "rhs"), environment)

		if reflect.TypeOf(lhs).Kind() == reflect.Float64 {
			lhs = int32(lhs.(float64))
		}

		if reflect.TypeOf(rhs).Kind() == reflect.Float64 {
			rhs = int32(rhs.(float64))
		}

		switch op := getNode(ast, "op").(string); op {
		case ADD:
			if reflect.TypeOf(lhs).Kind() == reflect.String || reflect.TypeOf(rhs).Kind() == reflect.String {
				if reflect.TypeOf(lhs).Kind() != reflect.String {
					lhs = fmt.Sprintf("%v", lhs)
				}

				if reflect.TypeOf(rhs).Kind() != reflect.String {
					rhs = fmt.Sprintf("%v", rhs)
				}

				return lhs.(string) + rhs.(string)
			} else {
				return lhs.(int32) + rhs.(int32)
			}
		case SUB:
			return lhs.(int32) - rhs.(int32)
		case MUL:
			return lhs.(int32) * rhs.(int32)
		case DIV:
			if rhs == 0.0 {
				panic("Division by zero") // return infinity?
			}
			return lhs.(int32) / rhs.(int32)
		case REM:
			return lhs.(int32) % rhs.(int32)
		case EQ:
			return fmt.Sprintf("%v", lhs) == fmt.Sprintf("%v", rhs)
		case NEQ:
			return fmt.Sprintf("%v", lhs) != fmt.Sprintf("%v", rhs)
		case LT:
			return lhs.(int32) < rhs.(int32)
		case GT:
			return lhs.(int32) > rhs.(int32)
		case LTE:
			return lhs.(int32) <= rhs.(int32)
		case GTE:
			return lhs.(int32) >= rhs.(int32)
		case AND:
			return lhs.(bool) && rhs.(bool)
		case OR:
			return lhs.(bool) || rhs.(bool)
		default:
			panic(fmt.Sprintf("Unknown operator: <%s>", op))
		}
	case BOOL:
		return getNode(ast, "value").(bool)
	case LET:
		copy_environment := copyEnvironment(environment)

		name := getNode(getNode(ast, "name"), "text").(string)
		value := Execute(getNode(ast, "value"), copy_environment)

		copy_environment[name] = value

		return Execute(getNode(ast, "next"), copy_environment)
	case VAR:
		name := getNode(ast, "text").(string)
		if value, ok := environment[name]; ok {
			return value
		} else {
			panic(fmt.Sprintf("Undefined variable: <%s>", name))
		}
	case CALL:
		copy_environment := copyEnvironment(environment)

		callee := Execute(getNode(ast, "callee"), environment)
		arguments := getNode(ast, "arguments").([]interface{})
		params := getNode(callee, "parameters").([]interface{})

		for i := 0; i < len(arguments); i++ {
			arguments[i] = Execute(arguments[i], environment)
		}

		for i := 0; i < len(params); i++ {
			if _, ok := params[i].(map[string]interface{}); ok {
				params[i] = getNode(params[i], "text")
			}
		}

		if len(arguments) != len(params) {
			panic(fmt.Sprintf("Expected %d arguments, but got %d", len(params), len(arguments)))
		}

		for i := 0; i < len(params); i++ {
			copy_environment[params[i].(string)] = arguments[i]
		}

		return Execute(getNode(callee, "value"), copy_environment)

	case FUNCTION:
		return ast
	case IF:
		condition := Execute(getNode(ast, "condition"), environment)
		if condition.(bool) {
			then := getNode(ast, "then")
			return Execute(then, environment)
		} else {
			else_ := getNode(ast, "otherwise")
			return Execute(else_, environment)
		}
	default:
		panic(fmt.Sprintf("Unknown node kind: <%s>", kind))
	}
}
