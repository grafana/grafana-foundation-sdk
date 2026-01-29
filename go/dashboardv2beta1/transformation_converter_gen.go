// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// TransformationConverter accepts a `Transformation` object and generates the Go code to build this object using builders.
func TransformationConverter(input TransformationKind) string {
	calls := []string{
		`dashboardv2beta1.NewTransformationBuilder()`,
	}
	var buffer strings.Builder
	if input.Kind != "" {

		buffer.WriteString(`Kind(`)
		arg0 := fmt.Sprintf("%#v", input.Kind)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.Id != "" {

		buffer.WriteString(`Id(`)
		arg0 := fmt.Sprintf("%#v", input.Spec.Id)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.Disabled != nil {

		buffer.WriteString(`Disabled(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.Disabled)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.Filter != nil {

		buffer.WriteString(`Filter(`)
		arg0 := cog.Dump(*input.Spec.Filter)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.Topic != nil {

		buffer.WriteString(`Topic(`)
		arg0 := cog.Dump(*input.Spec.Topic)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.Options != nil {

		buffer.WriteString(`Options(`)
		arg0 := cog.Dump(input.Spec.Options)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
