// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// BuilderQueryEditorWhereExpressionArrayConverter accepts a `BuilderQueryEditorWhereExpressionArray` object and generates the Go code to build this object using builders.
func BuilderQueryEditorWhereExpressionArrayConverter(input BuilderQueryEditorWhereExpressionArray) string {
	calls := []string{
		`azuremonitor.NewBuilderQueryEditorWhereExpressionArrayBuilder()`,
	}
	var buffer strings.Builder
	if input.Expressions != nil && len(input.Expressions) >= 1 {

		buffer.WriteString(`Expressions(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Expressions {
			tmpexpressionsarg1 := BuilderQueryEditorWhereExpressionConverter(arg1)
			tmparg0 = append(tmparg0, tmpexpressionsarg1)
		}
		arg0 := "[]cog.Builder[azuremonitor.BuilderQueryEditorWhereExpression]{" + strings.Join(tmparg0, ",\n") + "}"
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
