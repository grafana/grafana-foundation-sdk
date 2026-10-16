// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"strings"
)

// CellKindConverter accepts a `CellKind` object and generates the Go code to build this object using builders.
func CellKindConverter(input CellKind) string {
	calls := []string{
		`dashboardv2beta1.NewCellKindBuilder()`,
	}
	var buffer strings.Builder

	{
		buffer.WriteString(`Spec(`)
		arg0 := CellSpecConverter(input.Spec)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	return strings.Join(calls, ".\t\n")
}
