// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	"fmt"
	"strings"
)

// DateHistogramSettingsConverter accepts a `DateHistogramSettings` object and generates the Go code to build this object using builders.
func DateHistogramSettingsConverter(input DateHistogramSettings) string {
	calls := []string{
		`elasticsearch.NewDateHistogramSettingsBuilder()`,
	}
	var buffer strings.Builder
	if input.Interval != nil && *input.Interval != "" {

		buffer.WriteString(`Interval(`)
		arg0 := fmt.Sprintf("%#v", *input.Interval)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.MinDocCount != nil && *input.MinDocCount != "" {

		buffer.WriteString(`MinDocCount(`)
		arg0 := fmt.Sprintf("%#v", *input.MinDocCount)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.TrimEdges != nil && *input.TrimEdges != "" {

		buffer.WriteString(`TrimEdges(`)
		arg0 := fmt.Sprintf("%#v", *input.TrimEdges)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Offset != nil && *input.Offset != "" {

		buffer.WriteString(`Offset(`)
		arg0 := fmt.Sprintf("%#v", *input.Offset)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.TimeZone != nil && *input.TimeZone != "" {

		buffer.WriteString(`TimeZone(`)
		arg0 := fmt.Sprintf("%#v", *input.TimeZone)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
