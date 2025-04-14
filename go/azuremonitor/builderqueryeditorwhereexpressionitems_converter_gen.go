// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// BuilderQueryEditorWhereExpressionItemsConverter accepts a `BuilderQueryEditorWhereExpressionItems` object and generates the Go code to build this object using builders.
func BuilderQueryEditorWhereExpressionItemsConverter(input BuilderQueryEditorWhereExpressionItems) string {
	calls := []string{
		`azuremonitor.NewBuilderQueryEditorWhereExpressionItemsBuilder()`,
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
		buffer.WriteString(`Operator(`)
		arg0 := BuilderQueryEditorOperatorConverter(input.Operator)
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
