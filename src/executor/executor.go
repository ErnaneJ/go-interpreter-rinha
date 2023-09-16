package src

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var ast struct {
	Expression interface{} `json:"expression"`
}

func Execute(filePath string) (interface{}, error) {
	rawData, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("An error occurred while reading the file:", err)
		return nil, err
	}

	if err := json.Unmarshal(rawData, &ast); err != nil {
		fmt.Println("An error occurred while parsing JSON:", err)
		return nil, err
	}

	return ast.Expression, nil
}
