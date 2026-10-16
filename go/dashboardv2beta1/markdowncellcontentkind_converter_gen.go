// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"strings"
)

// MarkdownCellContentKindConverter accepts a `MarkdownCellContentKind` object and generates the Go code to build this object using builders.
func MarkdownCellContentKindConverter(input MarkdownCellContentKind) string {
	calls := []string{
		`dashboardv2beta1.NewMarkdownCellContentKindBuilder()`,
	}
	var buffer strings.Builder

	{
		buffer.WriteString(`Spec(`)
		arg0 := MarkdownCellContentSpecConverter(input.Spec)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	return strings.Join(calls, ".\t\n")
}
