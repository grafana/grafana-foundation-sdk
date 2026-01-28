// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"strings"
)

// SwitchVariableKindConverter accepts a `SwitchVariableKind` object and generates the Go code to build this object using builders.
func SwitchVariableKindConverter(input SwitchVariableKind) string {
	calls := []string{
		`dashboardv2beta1.NewSwitchVariableKindBuilder()`,
	}
	var buffer strings.Builder

	{
		buffer.WriteString(`Spec(`)
		arg0 := SwitchVariableSpecConverter(input.Spec)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	return strings.Join(calls, ".\t\n")
}
