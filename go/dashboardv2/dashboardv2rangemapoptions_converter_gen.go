// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2

import (
	"fmt"
	"strings"
)

// Dashboardv2RangeMapOptionsConverter accepts a `Dashboardv2RangeMapOptions` object and generates the Go code to build this object using builders.
func Dashboardv2RangeMapOptionsConverter(input Dashboardv2RangeMapOptions) string {
	calls := []string{
		`dashboardv2.NewDashboardv2RangeMapOptionsBuilder()`,
	}
	var buffer strings.Builder
	if input.From != nil {

		buffer.WriteString(`From(`)
		arg0 := fmt.Sprintf("%#v", *input.From)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.To != nil {

		buffer.WriteString(`To(`)
		arg0 := fmt.Sprintf("%#v", *input.To)
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
