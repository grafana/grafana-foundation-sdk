// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboard

import (
	"fmt"
	"strings"
)

// DashboardRangeMapOptionsConverter accepts a `DashboardRangeMapOptions` object and generates the Go code to build this object using builders.
func DashboardRangeMapOptionsConverter(input DashboardRangeMapOptions) string {
	calls := []string{
		`dashboard.NewDashboardRangeMapOptionsBuilder()`,
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
