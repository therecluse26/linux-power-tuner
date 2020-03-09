package utils

import "github.com/maja42/goval"

func EvaluateExpression(exp interface{}) (interface{}, error) {
	result, err := goval.NewEvaluator().Evaluate(exp.(string), nil, nil)
	return result, err
}
