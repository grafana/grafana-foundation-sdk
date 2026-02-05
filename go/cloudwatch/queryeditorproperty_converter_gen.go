// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package cloudwatch

import (
	"fmt"
	"strings"
)

// QueryEditorPropertyConverter accepts a `QueryEditorProperty` object and generates the Go code to build this object using builders.
func QueryEditorPropertyConverter(input QueryEditorProperty) string {
	calls := []string{
		`cloudwatch.NewQueryEditorPropertyBuilder()`,
	}
	var buffer strings.Builder
	if input.Name != nil && *input.Name != "" {

		buffer.WriteString(`Name(`)
		arg0 := fmt.Sprintf("%#v", *input.Name)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
