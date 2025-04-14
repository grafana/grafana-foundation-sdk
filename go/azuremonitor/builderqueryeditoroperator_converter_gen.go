// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	"fmt"
	"strings"
)

// BuilderQueryEditorOperatorConverter accepts a `BuilderQueryEditorOperator` object and generates the Go code to build this object using builders.
func BuilderQueryEditorOperatorConverter(input BuilderQueryEditorOperator) string {
	calls := []string{
		`azuremonitor.NewBuilderQueryEditorOperatorBuilder()`,
	}
	var buffer strings.Builder
	if input.Name != "" {

		buffer.WriteString(`Name(`)
		arg0 := fmt.Sprintf("%#v", input.Name)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Value != "" {

		buffer.WriteString(`Value(`)
		arg0 := fmt.Sprintf("%#v", input.Value)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.LabelValue != nil && *input.LabelValue != "" {

		buffer.WriteString(`LabelValue(`)
		arg0 := fmt.Sprintf("%#v", *input.LabelValue)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
