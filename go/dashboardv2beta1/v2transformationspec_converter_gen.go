// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// V2TransformationSpecConverter accepts a `V2TransformationSpec` object and generates the Go code to build this object using builders.
func V2TransformationSpecConverter(input V2TransformationSpec) string {
	calls := []string{
		`dashboardv2beta1.NewV2TransformationSpecBuilder()`,
	}
	var buffer strings.Builder
	if input.Disabled != nil {

		buffer.WriteString(`Disabled(`)
		arg0 := fmt.Sprintf("%#v", *input.Disabled)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Filter != nil {

		buffer.WriteString(`Filter(`)
		arg0 := cog.Dump(*input.Filter)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Topic != nil {

		buffer.WriteString(`Topic(`)
		arg0 := cog.Dump(*input.Topic)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Options != nil {

		buffer.WriteString(`Options(`)
		arg0 := cog.Dump(input.Options)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
