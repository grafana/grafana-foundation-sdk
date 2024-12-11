// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboard

import (
	"fmt"
	"strings"
)

// ValueMappingResultConverter accepts a `ValueMappingResult` object and generates the Go code to build this object using builders.
func ValueMappingResultConverter(input ValueMappingResult) string {
	calls := []string{
		`dashboard.NewValueMappingResultBuilder()`,
	}
	var buffer strings.Builder
	if input.Text != nil && *input.Text != "" {

		buffer.WriteString(`Text(`)
		arg0 := fmt.Sprintf("%#v", *input.Text)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Color != nil && *input.Color != "" {

		buffer.WriteString(`Color(`)
		arg0 := fmt.Sprintf("%#v", *input.Color)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Icon != nil && *input.Icon != "" {

		buffer.WriteString(`Icon(`)
		arg0 := fmt.Sprintf("%#v", *input.Icon)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Index != nil {

		buffer.WriteString(`Index(`)
		arg0 := fmt.Sprintf("%#v", *input.Index)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
