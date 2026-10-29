// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alertingv0alpha1

import (
	"fmt"
	"strings"
)

// PanelRefConverter accepts a `PanelRef` object and generates the Go code to build this object using builders.
func PanelRefConverter(input PanelRef) string {
	calls := []string{
		`alertingv0alpha1.NewPanelRefBuilder()`,
	}
	var buffer strings.Builder
	if input.DashboardUID != "" {

		buffer.WriteString(`DashboardUID(`)
		arg0 := fmt.Sprintf("%#v", input.DashboardUID)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	{
		buffer.WriteString(`PanelID(`)
		arg0 := fmt.Sprintf("%#v", input.PanelID)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	return strings.Join(calls, ".\t\n")
}
