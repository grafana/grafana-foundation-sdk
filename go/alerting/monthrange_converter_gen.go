// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alerting

import (
	"fmt"
	"strings"
)

// MonthRangeConverter accepts a `MonthRange` object and generates the Go code to build this object using builders.
func MonthRangeConverter(input MonthRange) string {
	calls := []string{
		`alerting.NewMonthRangeBuilder()`,
	}
	var buffer strings.Builder
	if input.Begin != nil {

		buffer.WriteString(`Begin(`)
		arg0 := fmt.Sprintf("%#v", *input.Begin)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.End != nil {

		buffer.WriteString(`End(`)
		arg0 := fmt.Sprintf("%#v", *input.End)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
