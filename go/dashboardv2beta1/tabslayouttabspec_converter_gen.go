// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// TabsLayoutTabSpecConverter accepts a `TabsLayoutTabSpec` object and generates the Go code to build this object using builders.
func TabsLayoutTabSpecConverter(input TabsLayoutTabSpec) string {
	calls := []string{
		`dashboardv2beta1.NewTabsLayoutTabSpecBuilder()`,
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

	{
		buffer.WriteString(`Layout(`)
		arg0 := cog.Dump(input.Layout)
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
		arg0 := TabRepeatOptionsConverter(*input.Repeat)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
