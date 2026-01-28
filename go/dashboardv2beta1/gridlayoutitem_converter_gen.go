// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"fmt"
	"strings"
)

// GridLayoutItemConverter accepts a `GridLayoutItem` object and generates the Go code to build this object using builders.
func GridLayoutItemConverter(input GridLayoutItemKind) string {
	calls := []string{
		`dashboardv2beta1.NewGridLayoutItemBuilder(` + fmt.Sprintf("%#v", input.Spec.Element.Name) + `)`,
	}
	var buffer strings.Builder

	{
		buffer.WriteString(`X(`)
		arg0 := fmt.Sprintf("%#v", input.Spec.X)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	{
		buffer.WriteString(`Y(`)
		arg0 := fmt.Sprintf("%#v", input.Spec.Y)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	{
		buffer.WriteString(`Width(`)
		arg0 := fmt.Sprintf("%#v", input.Spec.Width)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	{
		buffer.WriteString(`Height(`)
		arg0 := fmt.Sprintf("%#v", input.Spec.Height)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.Spec.Repeat != nil {

		buffer.WriteString(`Repeat(`)
		arg0 := RepeatOptionsConverter(*input.Spec.Repeat)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.Element.Name != "" {

		buffer.WriteString(`Element(`)
		arg0 := fmt.Sprintf("%#v", input.Spec.Element.Name)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
