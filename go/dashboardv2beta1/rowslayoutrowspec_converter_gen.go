// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// RowsLayoutRowSpecConverter accepts a `RowsLayoutRowSpec` object and generates the Go code to build this object using builders.
func RowsLayoutRowSpecConverter(input RowsLayoutRowSpec) string {
	calls := []string{
		`dashboardv2beta1.NewRowsLayoutRowSpecBuilder()`,
	}
	var buffer strings.Builder
	if input.Title != nil && *input.Title != "" {

		buffer.WriteString(`Title(`)
		arg0 := fmt.Sprintf("%#v", *input.Title)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Collapse != nil {

		buffer.WriteString(`Collapse(`)
		arg0 := fmt.Sprintf("%#v", *input.Collapse)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.HideHeader != nil {

		buffer.WriteString(`HideHeader(`)
		arg0 := fmt.Sprintf("%#v", *input.HideHeader)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.FillScreen != nil {

		buffer.WriteString(`FillScreen(`)
		arg0 := fmt.Sprintf("%#v", *input.FillScreen)
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
	if input.Repeat != nil {

		buffer.WriteString(`Repeat(`)
		arg0 := RowRepeatOptionsConverter(*input.Repeat)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	{
		buffer.WriteString(`Layout(`)
		arg0 := cog.Dump(input.Layout)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	return strings.Join(calls, ".\t\n")
}
