// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// BuilderQueryEditorColumnsExpressionConverter accepts a `BuilderQueryEditorColumnsExpression` object and generates the Go code to build this object using builders.
func BuilderQueryEditorColumnsExpressionConverter(input BuilderQueryEditorColumnsExpression) string {
	calls := []string{
		`azuremonitor.NewBuilderQueryEditorColumnsExpressionBuilder()`,
	}
	var buffer strings.Builder
	if input.Columns != nil && len(input.Columns) >= 1 {

		buffer.WriteString(`Columns(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Columns {
			tmpcolumnsarg1 := fmt.Sprintf("%#v", arg1)
			tmparg0 = append(tmparg0, tmpcolumnsarg1)
		}
		arg0 := "[]string{" + strings.Join(tmparg0, ",\n") + "}"
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
