// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboard

import (
	"fmt"
	"strings"
)

// DashboardActionStyleConverter accepts a `DashboardActionStyle` object and generates the Go code to build this object using builders.
func DashboardActionStyleConverter(input DashboardActionStyle) string {
	calls := []string{
		`dashboard.NewDashboardActionStyleBuilder()`,
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
