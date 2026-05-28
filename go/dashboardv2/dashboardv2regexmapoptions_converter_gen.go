// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2

import (
	"fmt"
	"strings"
)

// Dashboardv2RegexMapOptionsConverter accepts a `Dashboardv2RegexMapOptions` object and generates the Go code to build this object using builders.
func Dashboardv2RegexMapOptionsConverter(input Dashboardv2RegexMapOptions) string {
	calls := []string{
		`dashboardv2.NewDashboardv2RegexMapOptionsBuilder()`,
	}
	var buffer strings.Builder
	if input.Pattern != "" {

		buffer.WriteString(`Pattern(`)
		arg0 := fmt.Sprintf("%#v", input.Pattern)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	{
		buffer.WriteString(`Result(`)
		arg0 := ValueMappingResultConverter(input.Result)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	return strings.Join(calls, ".\t\n")
}
