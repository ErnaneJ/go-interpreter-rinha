package src

import "fmt"

var dataFromInterpreter interface{}

func getNode(ast interface{}, key string) interface{} {
	return ast.(map[string]interface{})[key]
}

func Interpret(astExpression interface{}) interface{} {
	switch kind := getNode(astExpression, "kind"); kind {
	case "Print":
		termoo := Interpret(getNode(astExpression, "value"))
		fmt.Print(termoo)
		return nil
	case "Str":
		return getNode(astExpression, "value").(string)
	case "Int":
		return getNode(astExpression, "value").(float64)
	default:
		fmt.Printf("Unknown node kind: <%s>", kind)
	}

	return nil
}
