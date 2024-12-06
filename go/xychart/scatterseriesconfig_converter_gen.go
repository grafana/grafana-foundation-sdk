// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package xychart

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
)

// ScatterSeriesConfigConverter accepts a `ScatterSeriesConfig` object and generates the Go code to build this object using builders.
func ScatterSeriesConfigConverter(input ScatterSeriesConfig) string {
	calls := []string{
		`xychart.NewScatterSeriesConfigBuilder()`,
	}
	var buffer strings.Builder
	if input.X != nil && *input.X != "" {

		buffer.WriteString(`X(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.X))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Y != nil && *input.Y != "" {

		buffer.WriteString(`Y(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Y))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Name != nil && *input.Name != "" {

		buffer.WriteString(`Name(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Name))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Show != nil {

		buffer.WriteString(`Show(`)
		arg0 := cog.Dump(*input.Show)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.PointSize != nil {

		buffer.WriteString(`PointSize(`)
		arg0 := common.ScaleDimensionConfigConverter(*input.PointSize)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.PointColor != nil {

		buffer.WriteString(`PointColor(`)
		arg0 := common.ColorDimensionConfigConverter(*input.PointColor)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.LineColor != nil {

		buffer.WriteString(`LineColor(`)
		arg0 := common.ColorDimensionConfigConverter(*input.LineColor)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.LineWidth != nil {

		buffer.WriteString(`LineWidth(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.LineWidth))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.LineStyle != nil {

		buffer.WriteString(`LineStyle(`)
		arg0 := common.LineStyleConverter(*input.LineStyle)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Label != nil {

		buffer.WriteString(`Label(`)
		arg0 := cog.Dump(*input.Label)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.HideFrom != nil {

		buffer.WriteString(`HideFrom(`)
		arg0 := common.HideSeriesConfigConverter(*input.HideFrom)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.AxisPlacement != nil {

		buffer.WriteString(`AxisPlacement(`)
		arg0 := cog.Dump(*input.AxisPlacement)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.AxisColorMode != nil {

		buffer.WriteString(`AxisColorMode(`)
		arg0 := cog.Dump(*input.AxisColorMode)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.AxisLabel != nil && *input.AxisLabel != "" {

		buffer.WriteString(`AxisLabel(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.AxisLabel))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.AxisWidth != nil {

		buffer.WriteString(`AxisWidth(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.AxisWidth))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.AxisSoftMin != nil {

		buffer.WriteString(`AxisSoftMin(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.AxisSoftMin))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.AxisSoftMax != nil {

		buffer.WriteString(`AxisSoftMax(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.AxisSoftMax))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.AxisGridShow != nil {

		buffer.WriteString(`AxisGridShow(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.AxisGridShow))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.ScaleDistribution != nil {

		buffer.WriteString(`ScaleDistribution(`)
		arg0 := common.ScaleDistributionConfigConverter(*input.ScaleDistribution)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.AxisCenteredZero != nil {

		buffer.WriteString(`AxisCenteredZero(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.AxisCenteredZero))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Frame != nil {

		buffer.WriteString(`Frame(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Frame))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.LabelValue != nil {

		buffer.WriteString(`LabelValue(`)
		arg0 := common.TextDimensionConfigConverter(*input.LabelValue)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.AxisBorderShow != nil {

		buffer.WriteString(`AxisBorderShow(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.AxisBorderShow))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
