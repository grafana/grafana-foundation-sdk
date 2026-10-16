// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"strings"
)

// V2QueryGroupKindConverter accepts a `V2QueryGroupKind` object and generates the Go code to build this object using builders.
func V2QueryGroupKindConverter(input V2QueryGroupKind) string {
	calls := []string{
		`dashboardv2beta1.NewV2QueryGroupKindBuilder()`,
	}
	var buffer strings.Builder

	{
		buffer.WriteString(`Spec(`)
		arg0 := V2QueryGroupSpecConverter(input.Spec)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	return strings.Join(calls, ".\t\n")
}
