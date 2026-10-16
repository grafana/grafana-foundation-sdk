// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"strings"
)

// V2PanelKindConverter accepts a `V2PanelKind` object and generates the Go code to build this object using builders.
func V2PanelKindConverter(input V2PanelKind) string {
	calls := []string{
		`dashboardv2beta1.NewV2PanelKindBuilder()`,
	}
	var buffer strings.Builder

	{
		buffer.WriteString(`Spec(`)
		arg0 := V2PanelSpecConverter(input.Spec)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	return strings.Join(calls, ".\t\n")
}
