// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"fmt"
	"strings"
)

// Dashboardv2beta1ActionStyleConverter accepts a `Dashboardv2beta1ActionStyle` object and generates the Go code to build this object using builders.
func Dashboardv2beta1ActionStyleConverter(input Dashboardv2beta1ActionStyle) string {
	calls := []string{
		`dashboardv2beta1.NewDashboardv2beta1ActionStyleBuilder()`,
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
