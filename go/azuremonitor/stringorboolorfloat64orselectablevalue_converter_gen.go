// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	"fmt"
	"strings"
)

// StringOrBoolOrFloat64OrSelectableValueConverter accepts a `StringOrBoolOrFloat64OrSelectableValue` object and generates the Go code to build this object using builders.
func StringOrBoolOrFloat64OrSelectableValueConverter(input StringOrBoolOrFloat64OrSelectableValue) string {
	calls := []string{
		`azuremonitor.NewStringOrBoolOrFloat64OrSelectableValueBuilder()`,
	}
	var buffer strings.Builder
	if input.String != nil && *input.String != "" {

		buffer.WriteString(`String(`)
		arg0 := fmt.Sprintf("%#v", *input.String)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Bool != nil {

		buffer.WriteString(`Bool(`)
		arg0 := fmt.Sprintf("%#v", *input.Bool)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Float64 != nil {

		buffer.WriteString(`Float64(`)
		arg0 := fmt.Sprintf("%#v", *input.Float64)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.SelectableValue != nil {

		buffer.WriteString(`SelectableValue(`)
		arg0 := SelectableValueConverter(*input.SelectableValue)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
