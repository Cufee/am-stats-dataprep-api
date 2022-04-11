package logic

import (
	"fmt"
)

type ExpressionTag string

const (
	SessionValue = "sessionValue"
	SessionOf    = "sessionOf"
	AllTimeValue = "allTimeValue"
	AllTimeOf    = "allTimeOf"
)

type Template struct {
	Expression string `json:"content"`
	Format     string `json:"format"`
}

func (t *Template) Evaluate(values Values) string {
	result, err := EvaluateExpression(t.Expression, values)
	if err != nil {
		return "invalid expression"
	}
	return fmt.Sprintf(t.Format, result)
}
