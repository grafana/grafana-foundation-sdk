// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package heatmap

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// HeatmapColorOptionsConverter accepts a `HeatmapColorOptions` object and generates the Go code to build this object using builders.
func HeatmapColorOptionsConverter(input HeatmapColorOptions) string {
	calls := []string{
		`heatmap.NewHeatmapColorOptionsBuilder()`,
	}
	var buffer strings.Builder
	if input.Mode != nil {

		buffer.WriteString(`Mode(`)
		arg0 := cog.Dump(*input.Mode)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Scheme != "" {

		buffer.WriteString(`Scheme(`)
		arg0 := fmt.Sprintf("%#v", input.Scheme)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Fill != "" {

		buffer.WriteString(`Fill(`)
		arg0 := fmt.Sprintf("%#v", input.Fill)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Scale != nil {

		buffer.WriteString(`Scale(`)
		arg0 := cog.Dump(*input.Scale)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	{
		buffer.WriteString(`Exponent(`)
		arg0 := fmt.Sprintf("%#v", input.Exponent)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	{
		buffer.WriteString(`Steps(`)
		arg0 := fmt.Sprintf("%#v", input.Steps)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	{
		buffer.WriteString(`Reverse(`)
		arg0 := fmt.Sprintf("%#v", input.Reverse)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.Min != nil {

		buffer.WriteString(`Min(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Min))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Max != nil {

		buffer.WriteString(`Max(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Max))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
