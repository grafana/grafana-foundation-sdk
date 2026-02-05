// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package bigquery

import (
	"fmt"
	"strings"
)

// QueryEditorFunctionParameterExpressionConverter accepts a `QueryEditorFunctionParameterExpression` object and generates the Go code to build this object using builders.
func QueryEditorFunctionParameterExpressionConverter(input QueryEditorFunctionParameterExpression) string {
	calls := []string{
		`bigquery.NewQueryEditorFunctionParameterExpressionBuilder()`,
	}
	var buffer strings.Builder
	if input.Name != nil && *input.Name != "" {

		buffer.WriteString(`Name(`)
		arg0 := fmt.Sprintf("%#v", *input.Name)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
