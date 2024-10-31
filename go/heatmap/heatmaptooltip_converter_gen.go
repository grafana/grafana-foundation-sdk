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
		buffer.WriteString(`Show(`)
		arg0 := fmt.Sprintf("%#v", input.Show)
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

	return strings.Join(calls, ".\t\n")
}
