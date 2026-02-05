// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package xychart

import (
	"strings"
)

// XychartXYSeriesConfigXConverter accepts a `XychartXYSeriesConfigX` object and generates the Go code to build this object using builders.
func XychartXYSeriesConfigXConverter(input XychartXYSeriesConfigX) string {
	calls := []string{
		`xychart.NewXychartXYSeriesConfigXBuilder()`,
	}
	var buffer strings.Builder

	{
		buffer.WriteString(`Matcher(`)
		arg0 := MatcherConfigConverter(input.Matcher)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	return strings.Join(calls, ".\t\n")
}
