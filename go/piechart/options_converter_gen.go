// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package piechart

import (
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
)

// OptionsConverter accepts a `Options` object and generates the Go code to build this object using builders.
func OptionsConverter(input Options) string {
	calls := []string{
		`piechart.NewOptionsBuilder()`,
	}
	var buffer strings.Builder

	{
		buffer.WriteString(`PieType(`)
		arg0 := cog.Dump(input.PieType)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.DisplayLabels != nil && len(input.DisplayLabels) >= 1 {

		buffer.WriteString(`DisplayLabels(`)
		tmparg0 := []string{}
		for _, arg1 := range input.DisplayLabels {
			tmpdisplayLabelsarg1 := cog.Dump(arg1)
			tmparg0 = append(tmparg0, tmpdisplayLabelsarg1)
		}
		arg0 := "[]piechart.PieChartLabels{" + strings.Join(tmparg0, ",\n") + "}"
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

	{
		buffer.WriteString(`ReduceOptions(`)
		arg0 := common.ReduceDataOptionsConverter(input.ReduceOptions)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.Text != nil {

		buffer.WriteString(`Text(`)
		arg0 := common.VizTextDisplayOptionsConverter(*input.Text)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	{
		buffer.WriteString(`Legend(`)
		arg0 := PieChartLegendOptionsConverter(input.Legend)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	{
		buffer.WriteString(`Orientation(`)
		arg0 := cog.Dump(input.Orientation)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	return strings.Join(calls, ".\t\n")
}
