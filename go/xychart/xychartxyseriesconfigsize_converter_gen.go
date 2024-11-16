// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package xychart

import (
	"strings"
)

// XychartXYSeriesConfigSizeConverter accepts a `XychartXYSeriesConfigSize` object and generates the Go code to build this object using builders.
func XychartXYSeriesConfigSizeConverter(input XychartXYSeriesConfigSize) string {
	calls := []string{
		`xychart.NewXychartXYSeriesConfigSizeBuilder()`,
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
