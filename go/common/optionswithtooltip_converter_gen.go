// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package common

import (
	"strings"
)

// OptionsWithTooltipConverter accepts a `OptionsWithTooltip` object and generates the Go code to build this object using builders.
func OptionsWithTooltipConverter(input OptionsWithTooltip) string {
	calls := []string{
		`common.NewOptionsWithTooltipBuilder()`,
	}
	var buffer strings.Builder

	{
		buffer.WriteString(`Tooltip(`)
		arg0 := VizTooltipOptionsConverter(input.Tooltip)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	return strings.Join(calls, ".\t\n")
}
