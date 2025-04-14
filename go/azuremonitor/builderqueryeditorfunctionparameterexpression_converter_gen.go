// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// BuilderQueryEditorFunctionParameterExpressionConverter accepts a `BuilderQueryEditorFunctionParameterExpression` object and generates the Go code to build this object using builders.
func BuilderQueryEditorFunctionParameterExpressionConverter(input BuilderQueryEditorFunctionParameterExpression) string {
	calls := []string{
		`azuremonitor.NewBuilderQueryEditorFunctionParameterExpressionBuilder()`,
	}
	var buffer strings.Builder
	if input.Value != "" {

		buffer.WriteString(`Value(`)
		arg0 := fmt.Sprintf("%#v", input.Value)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	{
		buffer.WriteString(`FieldType(`)
		arg0 := cog.Dump(input.FieldType)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	{
		buffer.WriteString(`Type(`)
		arg0 := cog.Dump(input.Type)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	return strings.Join(calls, ".\t\n")
}
