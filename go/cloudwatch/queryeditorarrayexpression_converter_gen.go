// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package cloudwatch

import (
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// QueryEditorArrayExpressionConverter accepts a `QueryEditorArrayExpression` object and generates the Go code to build this object using builders.
func QueryEditorArrayExpressionConverter(input QueryEditorArrayExpression) string {
	calls := []string{
		`cloudwatch.NewQueryEditorArrayExpressionBuilder()`,
	}
	var buffer strings.Builder

	{
		buffer.WriteString(`Type(`)
		arg0 := cog.Dump(input.Type)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.Expressions != nil && len(input.Expressions) >= 1 {

		buffer.WriteString(`Expressions(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Expressions {
			tmpexpressionsarg1 := cog.Dump(arg1)
			tmparg0 = append(tmparg0, tmpexpressionsarg1)
		}
		arg0 := "[]cloudwatch.QueryEditorExpression{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
