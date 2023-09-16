package main

import (
	"os"

	executor "github.com/ernanej/go-interpreter-rinha/v1/src/executor"
	interpreter "github.com/ernanej/go-interpreter-rinha/v1/src/interpreter"
)

func main() {
	if len(os.Args) < 2 {
		panic("To execute the desired program prediction. for example: 'go run print' to interpret the file './var/rinha/print.rinha.json'")
	}

	fileName := os.Args[1]

	astExpression, err := executor.Execute("./var/rinha/" + fileName + ".rinha.json")

	if err != nil {
		panic(err)
	}

	interpreter.Interpret(astExpression)
}
