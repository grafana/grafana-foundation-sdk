// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"fmt"
	"strings"
)

// MarkdownCellContentSpecConverter accepts a `MarkdownCellContentSpec` object and generates the Go code to build this object using builders.
func MarkdownCellContentSpecConverter(input MarkdownCellContentSpec) string {
	calls := []string{
		`dashboardv2beta1.NewMarkdownCellContentSpecBuilder()`,
	}
	var buffer strings.Builder
	if input.Text != "" {

		buffer.WriteString(`Text(`)
		arg0 := fmt.Sprintf("%#v", input.Text)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
