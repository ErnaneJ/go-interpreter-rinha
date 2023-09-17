package main

import (
	Executor "github.com/ernanej/go-interpreter-rinha/v1/src/executor"
	Interpreter "github.com/ernanej/go-interpreter-rinha/v1/src/interpreter"
)

func main() {
	astExpression, err := Executor.Execute("./var/rinha/source.rinha.json")

	if err != nil {
		panic(err)
	}

	environment := make(map[string]interface{})
	Interpreter.Execute(astExpression, environment)
}
