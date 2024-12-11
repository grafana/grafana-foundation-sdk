// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package xychart

import (
	"fmt"
	"strings"
)

// XychartXYSeriesConfigNameConverter accepts a `XychartXYSeriesConfigName` object and generates the Go code to build this object using builders.
func XychartXYSeriesConfigNameConverter(input XychartXYSeriesConfigName) string {
	calls := []string{
		`xychart.NewXychartXYSeriesConfigNameBuilder()`,
	}
	var buffer strings.Builder
	if input.Fixed != nil && *input.Fixed != "" {

		buffer.WriteString(`Fixed(`)
		arg0 := fmt.Sprintf("%#v", *input.Fixed)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
