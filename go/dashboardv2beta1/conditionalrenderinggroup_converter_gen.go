// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// ConditionalRenderingGroupConverter accepts a `ConditionalRenderingGroup` object and generates the Go code to build this object using builders.
func ConditionalRenderingGroupConverter(input ConditionalRenderingGroupKind) string {
	calls := []string{
		`dashboardv2beta1.NewConditionalRenderingGroupBuilder()`,
	}
	var buffer strings.Builder

	{
		buffer.WriteString(`Visibility(`)
		arg0 := cog.Dump(input.Spec.Visibility)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	{
		buffer.WriteString(`Condition(`)
		arg0 := cog.Dump(input.Spec.Condition)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.Spec.Items != nil && len(input.Spec.Items) >= 1 {

		buffer.WriteString(`Items(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Spec.Items {
			tmpitemsarg1 := cog.Dump(arg1)
			tmparg0 = append(tmparg0, tmpitemsarg1)
		}
		arg0 := "[]dashboardv2beta1.ConditionalRenderingVariableKindOrConditionalRenderingDataKindOrConditionalRenderingTimeRangeSizeKind{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
