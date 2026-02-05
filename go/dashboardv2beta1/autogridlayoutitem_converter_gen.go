// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"fmt"
	"strings"
)

// AutoGridLayoutItemConverter accepts a `AutoGridLayoutItem` object and generates the Go code to build this object using builders.
func AutoGridLayoutItemConverter(input AutoGridLayoutItemKind) string {
	calls := []string{
		`dashboardv2beta1.NewAutoGridLayoutItemBuilder(` + fmt.Sprintf("%#v", input.Spec.Element.Name) + `)`,
	}
	var buffer strings.Builder

	{
		buffer.WriteString(`Element(`)
		arg0 := ElementReferenceConverter(input.Spec.Element)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.Spec.Repeat != nil {

		buffer.WriteString(`Repeat(`)
		arg0 := AutoGridRepeatOptionsConverter(*input.Spec.Repeat)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.ConditionalRendering != nil {

		buffer.WriteString(`ConditionalRendering(`)
		arg0 := ConditionalRenderingGroupConverter(*input.Spec.ConditionalRendering)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.Element.Name != "" {

		buffer.WriteString(`Name(`)
		arg0 := fmt.Sprintf("%#v", input.Spec.Element.Name)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
