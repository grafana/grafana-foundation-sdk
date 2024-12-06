// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package heatmap

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// HeatmapTooltipConverter accepts a `HeatmapTooltip` object and generates the Go code to build this object using builders.
func HeatmapTooltipConverter(input HeatmapTooltip) string {
	calls := []string{
		`heatmap.NewHeatmapTooltipBuilder()`,
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

	if input.MaxHeight != nil {

		buffer.WriteString(`MaxHeight(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.MaxHeight))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.MaxWidth != nil {

		buffer.WriteString(`MaxWidth(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.MaxWidth))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.YHistogram != nil {

		buffer.WriteString(`YHistogram(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.YHistogram))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.ShowColorScale != nil {

		buffer.WriteString(`ShowColorScale(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.ShowColorScale))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
