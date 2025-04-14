// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// BuilderQueryEditorPropertyExpressionConverter accepts a `BuilderQueryEditorPropertyExpression` object and generates the Go code to build this object using builders.
func BuilderQueryEditorPropertyExpressionConverter(input BuilderQueryEditorPropertyExpression) string {
	calls := []string{
		`azuremonitor.NewBuilderQueryEditorPropertyExpressionBuilder()`,
	}
	var buffer strings.Builder

	{
		buffer.WriteString(`Property(`)
		arg0 := BuilderQueryEditorPropertyConverter(input.Property)
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
