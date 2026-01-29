// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"fmt"
	"strings"
)

// Dashboardv2beta1RegexMapOptionsConverter accepts a `Dashboardv2beta1RegexMapOptions` object and generates the Go code to build this object using builders.
func Dashboardv2beta1RegexMapOptionsConverter(input Dashboardv2beta1RegexMapOptions) string {
	calls := []string{
		`dashboardv2beta1.NewDashboardv2beta1RegexMapOptionsBuilder()`,
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
