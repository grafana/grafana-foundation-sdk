// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package geomap

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// ControlsOptionsConverter accepts a `ControlsOptions` object and generates the Go code to build this object using builders.
func ControlsOptionsConverter(input ControlsOptions) string {
	calls := []string{
		`geomap.NewControlsOptionsBuilder()`,
	}
	var buffer strings.Builder
	if input.ShowZoom != nil {

		buffer.WriteString(`ShowZoom(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.ShowZoom))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.MouseWheelZoom != nil {

		buffer.WriteString(`MouseWheelZoom(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.MouseWheelZoom))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.ShowAttribution != nil {

		buffer.WriteString(`ShowAttribution(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.ShowAttribution))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.ShowScale != nil {

		buffer.WriteString(`ShowScale(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.ShowScale))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.ShowDebug != nil {

		buffer.WriteString(`ShowDebug(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.ShowDebug))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.ShowMeasure != nil {

		buffer.WriteString(`ShowMeasure(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.ShowMeasure))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
