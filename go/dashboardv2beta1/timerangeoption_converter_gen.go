// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"fmt"
	"strings"
)

// TimeRangeOptionConverter accepts a `TimeRangeOption` object and generates the Go code to build this object using builders.
func TimeRangeOptionConverter(input TimeRangeOption) string {
	calls := []string{
		`dashboardv2beta1.NewTimeRangeOptionBuilder()`,
	}
	var buffer strings.Builder
	if input.Display != "" && input.Display != "Last 6 hours" {

		buffer.WriteString(`Display(`)
		arg0 := fmt.Sprintf("%#v", input.Display)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.From != "" && input.From != "now-6h" {

		buffer.WriteString(`From(`)
		arg0 := fmt.Sprintf("%#v", input.From)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.To != "" && input.To != "now" {

		buffer.WriteString(`To(`)
		arg0 := fmt.Sprintf("%#v", input.To)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
