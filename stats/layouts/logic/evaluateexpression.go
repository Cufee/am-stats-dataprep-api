package logic

import (
	"fmt"

	"github.com/Knetic/govaluate"
)

func EvaluateExpression(expression string, values Values) (string, error) {
	expr, err := govaluate.NewEvaluableExpression(expression)
	if err != nil {
		return "", err
	}

	result, err := expr.Evaluate(values)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v", result), nil
}
