// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package xychart

import (
	"fmt"
	"strings"
)

// XychartFieldConfigPointSizeConverter accepts a `XychartFieldConfigPointSize` object and generates the Go code to build this object using builders.
func XychartFieldConfigPointSizeConverter(input XychartFieldConfigPointSize) string {
	calls := []string{
		`xychart.NewXychartFieldConfigPointSizeBuilder()`,
	}
	var buffer strings.Builder
	if input.Fixed != nil {

		buffer.WriteString(`Fixed(`)
		arg0 := fmt.Sprintf("%#v", *input.Fixed)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Min != nil {

		buffer.WriteString(`Min(`)
		arg0 := fmt.Sprintf("%#v", *input.Min)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Max != nil {

		buffer.WriteString(`Max(`)
		arg0 := fmt.Sprintf("%#v", *input.Max)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
