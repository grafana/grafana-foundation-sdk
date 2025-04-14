// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	"fmt"
	"strings"
)

// BuilderQueryEditorReduceExpressionConverter accepts a `BuilderQueryEditorReduceExpression` object and generates the Go code to build this object using builders.
func BuilderQueryEditorReduceExpressionConverter(input BuilderQueryEditorReduceExpression) string {
	calls := []string{
		`azuremonitor.NewBuilderQueryEditorReduceExpressionBuilder()`,
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
	if input.Reduce != nil {

		buffer.WriteString(`Reduce(`)
		arg0 := BuilderQueryEditorPropertyConverter(*input.Reduce)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Parameters != nil && len(input.Parameters) >= 1 {

		buffer.WriteString(`Parameters(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Parameters {
			tmpparametersarg1 := BuilderQueryEditorFunctionParameterExpressionConverter(arg1)
			tmparg0 = append(tmparg0, tmpparametersarg1)
		}
		arg0 := "[]cog.Builder[azuremonitor.BuilderQueryEditorFunctionParameterExpression]{" + strings.Join(tmparg0, ",\n") + "}"
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

	return strings.Join(calls, ".\t\n")
}
