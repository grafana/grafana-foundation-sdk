// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package bigquery

import (
	"strings"
)

// QueryEditorPropertyExpressionConverter accepts a `QueryEditorPropertyExpression` object and generates the Go code to build this object using builders.
func QueryEditorPropertyExpressionConverter(input QueryEditorPropertyExpression) string {
	calls := []string{
		`bigquery.NewQueryEditorPropertyExpressionBuilder()`,
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
