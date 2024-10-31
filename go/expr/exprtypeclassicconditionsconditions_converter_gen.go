// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package expr

import (
	"strings"
)

// ExprTypeClassicConditionsConditionsConverter accepts a `ExprTypeClassicConditionsConditions` object and generates the Go code to build this object using builders.
func ExprTypeClassicConditionsConditionsConverter(input ExprTypeClassicConditionsConditions) string {
	calls := []string{
		`expr.NewExprTypeClassicConditionsConditionsBuilder()`,
	}
	var buffer strings.Builder

	{
		buffer.WriteString(`Evaluator(`)
		arg0 := ExprTypeClassicConditionsConditionsEvaluatorConverter(input.Evaluator)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	{
		buffer.WriteString(`Operator(`)
		arg0 := ExprTypeClassicConditionsConditionsOperatorConverter(input.Operator)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	{
		buffer.WriteString(`Query(`)
		arg0 := ExprTypeClassicConditionsConditionsQueryConverter(input.Query)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	{
		buffer.WriteString(`Reducer(`)
		arg0 := ExprTypeClassicConditionsConditionsReducerConverter(input.Reducer)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	return strings.Join(calls, ".\t\n")
}
