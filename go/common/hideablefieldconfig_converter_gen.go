// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package common

import (
	"strings"
)

// HideableFieldConfigConverter accepts a `HideableFieldConfig` object and generates the Go code to build this object using builders.
func HideableFieldConfigConverter(input HideableFieldConfig) string {
	calls := []string{
		`common.NewHideableFieldConfigBuilder()`,
	}
	var buffer strings.Builder
	if input.HideFrom != nil {

		buffer.WriteString(`HideFrom(`)
		arg0 := HideSeriesConfigConverter(*input.HideFrom)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
