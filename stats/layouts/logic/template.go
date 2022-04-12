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
	String       = "string"
)

type Template struct {
	Expression string                   `json:"content"`
	Format     string                   `json:"format"`
	Parse      func(string) interface{} `json:"-"`
}

func (t *Template) Evaluate(values Values) string {
	result, err := EvaluateExpression(t.Expression, values)
	if err != nil {
		return "invalid expression"
	}
	if t.Parse != nil {
		return fmt.Sprintf(t.Format, t.Parse(result))
	}
	return fmt.Sprintf(t.Format, result)
}
