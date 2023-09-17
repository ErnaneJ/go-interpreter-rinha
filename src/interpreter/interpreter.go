package src

import (
	"fmt"
	"reflect"
)

func getNode(ast interface{}, key string) interface{} {
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
		value := getNode(ast, "value")
		if getNode(value, "kind") == nil {
			term := Execute(value, environment)
			fmt.Println(term)
		} else {
			tuple := Execute(value, environment)
			first := Execute(getNode(tuple, "first"), environment)
			second := Execute(getNode(tuple, "second"), environment)

			fmt.Println("(", first, ", ", second, ")")
		}
		return nil
	case STR, BOOL, INT:
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
	case LET:
		new_environment := copyEnvironment(environment)

		name := getNode(getNode(ast, "name"), "text").(string)
		value := Execute(getNode(ast, "value"), environment)

		new_environment[name] = value

		return Execute(getNode(ast, "next"), new_environment)
	case VAR:
		name := getNode(ast, "text").(string)
		if value, ok := environment[name]; ok {
			if _, ok := value.(map[string]interface{}); ok {
				return Execute(value, environment)
			} else {
				return value
			}
		} else {
			panic(fmt.Sprintf("Undefined variable: <%s>", name))
		}
	case CALL:
		callee := Execute(getNode(ast, "callee"), environment)

		if getNode(callee, "kind") != FUNCTION {
			return nil
		}

		arguments := getNode(ast, "arguments").([]interface{})
		params := getNode(callee, "parameters").([]interface{})

		for i := 0; i < len(arguments); i++ {
			if _, ok := arguments[i].(map[string]interface{}); ok {
				if getNode(arguments[i], "kind") == nil {
					arguments[i] = getNode(arguments[i], "value")
				} else {
					arguments[i] = Execute(arguments[i], environment)
				}
			}
		}

		for i := 0; i < len(params); i++ {
			if _, ok := params[i].(map[string]interface{}); ok {
				if getNode(params[i], "kind") == nil {
					params[i] = getNode(params[i], "text")
				} else {
					params[i] = Execute(params[i], environment)
				}
			}
		}

		if len(arguments) != len(params) {
			panic(fmt.Sprintf("Expected %d arguments, but got %d", len(params), len(arguments)))
		}

		new_environment := copyEnvironment(environment)

		for i := 0; i < len(params); i++ {
			new_environment[params[i].(string)] = arguments[i]
		}

		return Execute(getNode(callee, "value"), new_environment)
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
	case TUPLE:
		return ast
	case FIRST:
		top := getNode(ast, "value")
		if _, ok := top.(map[string]interface{}); ok {
			return Execute(getNode(Execute(top, environment), "first"), environment)
		} else {
			return Execute(getNode(top, "first"), environment)
		}
	case SECOND:
		top := getNode(ast, "value")
		if _, ok := top.(map[string]interface{}); ok {
			return Execute(getNode(Execute(top, environment), "second"), environment)
		} else {
			return Execute(getNode(top, "second"), environment)
		}
	default:
		panic(fmt.Sprintf("Unknown node kind: <%s>", kind))
	}
}
