package src

import "fmt"

func getNode(ast interface{}, key string) interface{} {
	return ast.(map[string]interface{})[key]
}

func Execute(astExpression interface{}, environment map[string]interface{}) interface{} {
	switch kind := getNode(astExpression, "kind"); kind {
	case PRINT:
		term := Execute(getNode(astExpression, "value"), environment)
		fmt.Println(term)
		return nil
	case STR:
		return getNode(astExpression, "value").(string)
	case INT:
		return getNode(astExpression, "value").(float64)
	case PARAMETER:
		return astExpression
	case BOOL:
		return getNode(astExpression, "value").(bool)
	case LET:
		name := getNode(getNode(astExpression, "name"), "text").(string)
		value := Execute(getNode(astExpression, "value"), environment)
		environment[name] = value
		Execute(getNode(astExpression, "next"), environment)

	case VAR:
		name := getNode(astExpression, "text").(string)
		if value, ok := environment[name]; ok {
			return value
		} else {
			panic(fmt.Sprintf("Undefined variable: <%s>", name))
		}
	case CALL:
		callee := Execute(getNode(astExpression, "callee"), environment)
		fmt.Print(callee)
		arguments := getNode(astExpression, "arguments").([]interface{})
		for i := 0; i < len(arguments); i++ {
			arguments[i] = Execute(arguments[i], environment)
		}
		fmt.Print(arguments)
	case IF:
		condition := Execute(getNode(astExpression, "condition"), environment)
		if condition.(bool) {
			then := getNode(astExpression, "then")
			return Execute(then, environment)
		} else {
			else_ := getNode(astExpression, "otherwise")
			return Execute(else_, environment)
		}
	default:
		fmt.Printf("Unknown node kind: <%s>", kind)
	}

	return nil
}
