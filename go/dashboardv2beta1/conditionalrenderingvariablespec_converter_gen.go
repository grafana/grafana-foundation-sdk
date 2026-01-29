// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// ConditionalRenderingVariableSpecConverter accepts a `ConditionalRenderingVariableSpec` object and generates the Go code to build this object using builders.
func ConditionalRenderingVariableSpecConverter(input ConditionalRenderingVariableSpec) string {
	calls := []string{
		`dashboardv2beta1.NewConditionalRenderingVariableSpecBuilder()`,
	}
	var buffer strings.Builder
	if input.Variable != "" {

		buffer.WriteString(`Variable(`)
		arg0 := fmt.Sprintf("%#v", input.Variable)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	{
		buffer.WriteString(`Operator(`)
		arg0 := cog.Dump(input.Operator)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.Value != "" {

		buffer.WriteString(`Value(`)
		arg0 := fmt.Sprintf("%#v", input.Value)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
