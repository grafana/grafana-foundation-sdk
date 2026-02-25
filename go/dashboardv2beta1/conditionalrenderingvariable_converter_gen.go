// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// ConditionalRenderingVariableConverter accepts a `ConditionalRenderingVariable` object and generates the Go code to build this object using builders.
func ConditionalRenderingVariableConverter(input ConditionalRenderingVariableKind) string {
	calls := []string{
		`dashboardv2beta1.NewConditionalRenderingVariableBuilder()`,
	}
	var buffer strings.Builder
	if input.Spec.Variable != "" {

		buffer.WriteString(`Variable(`)
		arg0 := fmt.Sprintf("%#v", input.Spec.Variable)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	{
		buffer.WriteString(`Operator(`)
		arg0 := cog.Dump(input.Spec.Operator)
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
