// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package cloudwatch

import (
	"strings"
)

// QueryEditorGroupByExpressionConverter accepts a `QueryEditorGroupByExpression` object and generates the Go code to build this object using builders.
func QueryEditorGroupByExpressionConverter(input QueryEditorGroupByExpression) string {
	calls := []string{
		`cloudwatch.NewQueryEditorGroupByExpressionBuilder()`,
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

	return strings.Join(calls, ".\t\n")
}
