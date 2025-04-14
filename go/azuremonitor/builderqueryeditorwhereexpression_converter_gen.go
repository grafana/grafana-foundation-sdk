// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// BuilderQueryEditorWhereExpressionConverter accepts a `BuilderQueryEditorWhereExpression` object and generates the Go code to build this object using builders.
func BuilderQueryEditorWhereExpressionConverter(input BuilderQueryEditorWhereExpression) string {
	calls := []string{
		`azuremonitor.NewBuilderQueryEditorWhereExpressionBuilder()`,
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
			tmpexpressionsarg1 := BuilderQueryEditorWhereExpressionItemsConverter(arg1)
			tmparg0 = append(tmparg0, tmpexpressionsarg1)
		}
		arg0 := "[]cog.Builder[azuremonitor.BuilderQueryEditorWhereExpressionItems]{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
