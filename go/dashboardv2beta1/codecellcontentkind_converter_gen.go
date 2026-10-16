// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"strings"
)

// CodeCellContentKindConverter accepts a `CodeCellContentKind` object and generates the Go code to build this object using builders.
func CodeCellContentKindConverter(input CodeCellContentKind) string {
	calls := []string{
		`dashboardv2beta1.NewCodeCellContentKindBuilder()`,
	}
	var buffer strings.Builder

	{
		buffer.WriteString(`Spec(`)
		arg0 := CodeCellContentSpecConverter(input.Spec)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	return strings.Join(calls, ".\t\n")
}
