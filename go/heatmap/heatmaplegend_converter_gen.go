// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package heatmap

import (
	"fmt"
	"strings"
)

// HeatmapLegendConverter accepts a `HeatmapLegend` object and generates the Go code to build this object using builders.
func HeatmapLegendConverter(input HeatmapLegend) string {
	calls := []string{
		`heatmap.NewHeatmapLegendBuilder()`,
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

	return strings.Join(calls, ".\t\n")
}
