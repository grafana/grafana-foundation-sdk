// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package xychart

import (
	"strings"
)

// XYSeriesConfigConverter accepts a `XYSeriesConfig` object and generates the Go code to build this object using builders.
func XYSeriesConfigConverter(input XYSeriesConfig) string {
	calls := []string{
		`xychart.NewXYSeriesConfigBuilder()`,
	}
	var buffer strings.Builder
	if input.Name != nil {

		buffer.WriteString(`Name(`)
		arg0 := XychartXYSeriesConfigNameConverter(*input.Name)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Frame != nil {

		buffer.WriteString(`Frame(`)
		arg0 := XychartXYSeriesConfigFrameConverter(*input.Frame)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.X != nil {

		buffer.WriteString(`X(`)
		arg0 := XychartXYSeriesConfigXConverter(*input.X)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Y != nil {

		buffer.WriteString(`Y(`)
		arg0 := XychartXYSeriesConfigYConverter(*input.Y)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Color != nil {

		buffer.WriteString(`Color(`)
		arg0 := XychartXYSeriesConfigColorConverter(*input.Color)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Size != nil {

		buffer.WriteString(`Size(`)
		arg0 := XychartXYSeriesConfigSizeConverter(*input.Size)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
