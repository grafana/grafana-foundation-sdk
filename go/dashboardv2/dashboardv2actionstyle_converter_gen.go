// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2

import (
	"fmt"
	"strings"
)

// Dashboardv2ActionStyleConverter accepts a `Dashboardv2ActionStyle` object and generates the Go code to build this object using builders.
func Dashboardv2ActionStyleConverter(input Dashboardv2ActionStyle) string {
	calls := []string{
		`dashboardv2.NewDashboardv2ActionStyleBuilder()`,
	}
	var buffer strings.Builder
	if input.BackgroundColor != nil && *input.BackgroundColor != "" {

		buffer.WriteString(`BackgroundColor(`)
		arg0 := fmt.Sprintf("%#v", *input.BackgroundColor)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
