// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"strings"
)

// AutoGridLayoutItemSpecConverter accepts a `AutoGridLayoutItemSpec` object and generates the Go code to build this object using builders.
func AutoGridLayoutItemSpecConverter(input AutoGridLayoutItemSpec) string {
	calls := []string{
		`dashboardv2beta1.NewAutoGridLayoutItemSpecBuilder()`,
	}
	var buffer strings.Builder

	{
		buffer.WriteString(`Element(`)
		arg0 := ElementReferenceConverter(input.Element)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.Repeat != nil {

		buffer.WriteString(`Repeat(`)
		arg0 := AutoGridRepeatOptionsConverter(*input.Repeat)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.ConditionalRendering != nil {

		buffer.WriteString(`ConditionalRendering(`)
		arg0 := ConditionalRenderingGroupConverter(*input.ConditionalRendering)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
