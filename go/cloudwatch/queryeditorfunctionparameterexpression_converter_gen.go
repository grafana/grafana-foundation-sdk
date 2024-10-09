// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package cloudwatch

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

func QueryEditorFunctionParameterExpressionConverter(input QueryEditorFunctionParameterExpression) string {
	calls := []string{
		`cloudwatch.NewQueryEditorFunctionParameterExpressionBuilder()`,
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

	return strings.Join(calls, ".\t\n")
}
