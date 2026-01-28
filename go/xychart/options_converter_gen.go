// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package xychart

import (
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
)

// OptionsConverter accepts a `Options` object and generates the Go code to build this object using builders.
func OptionsConverter(input Options) string {
	calls := []string{
		`xychart.NewOptionsBuilder()`,
	}
	var buffer strings.Builder
	if input.SeriesMapping != nil {

		buffer.WriteString(`SeriesMapping(`)
		arg0 := cog.Dump(*input.SeriesMapping)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	{
		buffer.WriteString(`Dims(`)
		arg0 := XYDimensionConfigConverter(input.Dims)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	{
		buffer.WriteString(`Legend(`)
		arg0 := common.VizLegendOptionsConverter(input.Legend)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	{
		buffer.WriteString(`Tooltip(`)
		arg0 := common.VizTooltipOptionsConverter(input.Tooltip)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.Series != nil && len(input.Series) >= 1 {

		buffer.WriteString(`Series(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Series {
			tmpseriesarg1 := ScatterSeriesConfigConverter(arg1)
			tmparg0 = append(tmparg0, tmpseriesarg1)
		}
		arg0 := "[]cog.Builder[xychart.ScatterSeriesConfig]{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
