// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package geomap

import (
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

func TooltipOptionsConverter(input TooltipOptions) string {
	calls := []string{
		`geomap.NewTooltipOptionsBuilder()`,
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

	return strings.Join(calls, ".\t\n")
}
