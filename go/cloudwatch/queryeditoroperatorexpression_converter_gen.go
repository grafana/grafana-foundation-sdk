// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package cloudwatch

import (
	"strings"
)

// QueryEditorOperatorExpressionConverter accepts a `QueryEditorOperatorExpression` object and generates the Go code to build this object using builders.
func QueryEditorOperatorExpressionConverter(input QueryEditorOperatorExpression) string {
	calls := []string{
		`cloudwatch.NewQueryEditorOperatorExpressionBuilder()`,
	}
	var buffer strings.Builder

	{
		buffer.WriteString(`Property(`)
		arg0 := QueryEditorPropertyConverter(input.Property)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	{
		buffer.WriteString(`Operator(`)
		arg0 := QueryEditorOperatorConverter(input.Operator)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	return strings.Join(calls, ".\t\n")
}
