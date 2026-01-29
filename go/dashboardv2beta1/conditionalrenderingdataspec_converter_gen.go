// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"fmt"
	"strings"
)

// ConditionalRenderingDataSpecConverter accepts a `ConditionalRenderingDataSpec` object and generates the Go code to build this object using builders.
func ConditionalRenderingDataSpecConverter(input ConditionalRenderingDataSpec) string {
	calls := []string{
		`dashboardv2beta1.NewConditionalRenderingDataSpecBuilder()`,
	}
	var buffer strings.Builder

	{
		buffer.WriteString(`Value(`)
		arg0 := fmt.Sprintf("%#v", input.Value)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	return strings.Join(calls, ".\t\n")
}
