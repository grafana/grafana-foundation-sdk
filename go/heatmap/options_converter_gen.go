// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package heatmap

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
)

// OptionsConverter accepts a `Options` object and generates the Go code to build this object using builders.
func OptionsConverter(input Options) string {
	calls := []string{
		`heatmap.NewOptionsBuilder()`,
	}
	var buffer strings.Builder
	if input.Calculate != nil && *input.Calculate != false {

		buffer.WriteString(`Calculate(`)
		arg0 := fmt.Sprintf("%#v", *input.Calculate)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Calculation != nil {

		buffer.WriteString(`Calculation(`)
		arg0 := common.HeatmapCalculationOptionsConverter(*input.Calculation)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	{
		buffer.WriteString(`Color(`)
		arg0 := HeatmapColorOptionsConverter(input.Color)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.FilterValues != nil {

		buffer.WriteString(`FilterValues(`)
		arg0 := FilterValueRangeConverter(*input.FilterValues)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.RowsFrame != nil {

		buffer.WriteString(`RowsFrame(`)
		arg0 := RowsHeatmapOptionsConverter(*input.RowsFrame)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	{
		buffer.WriteString(`ShowValue(`)
		arg0 := cog.Dump(input.ShowValue)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.CellGap != nil && *input.CellGap != 1 {

		buffer.WriteString(`CellGap(`)
		arg0 := fmt.Sprintf("%#v", *input.CellGap)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.CellRadius != nil {

		buffer.WriteString(`CellRadius(`)
		arg0 := fmt.Sprintf("%#v", *input.CellRadius)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.CellValues != nil {

		buffer.WriteString(`CellValues(`)
		arg0 := CellValuesConverter(*input.CellValues)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	{
		buffer.WriteString(`YAxis(`)
		arg0 := YAxisConfigConverter(input.YAxis)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	{
		buffer.WriteString(`Legend(`)
		arg0 := HeatmapLegendConverter(input.Legend)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	{
		buffer.WriteString(`Tooltip(`)
		arg0 := HeatmapTooltipConverter(input.Tooltip)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	{
		buffer.WriteString(`Exemplars(`)
		arg0 := ExemplarConfigConverter(input.Exemplars)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.SelectionMode != nil {

		buffer.WriteString(`SelectionMode(`)
		arg0 := cog.Dump(*input.SelectionMode)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
