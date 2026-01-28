// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"fmt"
	"strings"
)

// ConditionalRenderingTimeRangeSizeConverter accepts a `ConditionalRenderingTimeRangeSize` object and generates the Go code to build this object using builders.
func ConditionalRenderingTimeRangeSizeConverter(input ConditionalRenderingTimeRangeSizeKind) string {
	calls := []string{
		`dashboardv2beta1.NewConditionalRenderingTimeRangeSizeBuilder()`,
	}
	var buffer strings.Builder

	{
		buffer.WriteString(`Spec(`)
		arg0 := ConditionalRenderingTimeRangeSizeSpecConverter(input.Spec)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.Spec.Value != "" {

		buffer.WriteString(`Value(`)
		arg0 := fmt.Sprintf("%#v", input.Spec.Value)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
