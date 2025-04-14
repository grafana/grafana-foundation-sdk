// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// BuilderQueryEditorGroupByExpressionArrayConverter accepts a `BuilderQueryEditorGroupByExpressionArray` object and generates the Go code to build this object using builders.
func BuilderQueryEditorGroupByExpressionArrayConverter(input BuilderQueryEditorGroupByExpressionArray) string {
	calls := []string{
		`azuremonitor.NewBuilderQueryEditorGroupByExpressionArrayBuilder()`,
	}
	var buffer strings.Builder
	if input.Expressions != nil && len(input.Expressions) >= 1 {

		buffer.WriteString(`Expressions(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Expressions {
			tmpexpressionsarg1 := BuilderQueryEditorGroupByExpressionConverter(arg1)
			tmparg0 = append(tmparg0, tmpexpressionsarg1)
		}
		arg0 := "[]cog.Builder[azuremonitor.BuilderQueryEditorGroupByExpression]{" + strings.Join(tmparg0, ",\n") + "}"
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
