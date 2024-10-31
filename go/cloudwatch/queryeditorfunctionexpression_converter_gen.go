// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package cloudwatch

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// QueryEditorFunctionExpressionConverter accepts a `QueryEditorFunctionExpression` object and generates the Go code to build this object using builders.
func QueryEditorFunctionExpressionConverter(input QueryEditorFunctionExpression) string {
	calls := []string{
		`cloudwatch.NewQueryEditorFunctionExpressionBuilder()`,
	}
	var buffer strings.Builder
	if input.Name != nil && *input.Name != "" {

		buffer.WriteString(`Name(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Name))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Parameters != nil && len(input.Parameters) >= 1 {

		buffer.WriteString(`Parameters(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Parameters {
			tmpparametersarg1 := QueryEditorFunctionParameterExpressionConverter(arg1)
			tmparg0 = append(tmparg0, tmpparametersarg1)
		}
		arg0 := "[]cog.Builder[cloudwatch.QueryEditorFunctionParameterExpression]{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
