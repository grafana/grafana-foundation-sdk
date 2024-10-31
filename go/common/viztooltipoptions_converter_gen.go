// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package common

import (
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// VizTooltipOptionsConverter accepts a `VizTooltipOptions` object and generates the Go code to build this object using builders.
func VizTooltipOptionsConverter(input VizTooltipOptions) string {
	calls := []string{
		`common.NewVizTooltipOptionsBuilder()`,
	}
	var buffer strings.Builder

	{
		buffer.WriteString(`Mode(`)
		arg0 := cog.Dump(input.Mode)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	{
		buffer.WriteString(`Sort(`)
		arg0 := cog.Dump(input.Sort)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	return strings.Join(calls, ".\t\n")
}
