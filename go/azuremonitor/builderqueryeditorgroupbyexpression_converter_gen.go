// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// BuilderQueryEditorGroupByExpressionConverter accepts a `BuilderQueryEditorGroupByExpression` object and generates the Go code to build this object using builders.
func BuilderQueryEditorGroupByExpressionConverter(input BuilderQueryEditorGroupByExpression) string {
	calls := []string{
		`azuremonitor.NewBuilderQueryEditorGroupByExpressionBuilder()`,
	}
	var buffer strings.Builder
	if input.Property != nil {

		buffer.WriteString(`Property(`)
		arg0 := BuilderQueryEditorPropertyConverter(*input.Property)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Interval != nil {

		buffer.WriteString(`Interval(`)
		arg0 := BuilderQueryEditorPropertyConverter(*input.Interval)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Focus != nil {

		buffer.WriteString(`Focus(`)
		arg0 := fmt.Sprintf("%#v", *input.Focus)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Type != nil {

		buffer.WriteString(`Type(`)
		arg0 := cog.Dump(*input.Type)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
