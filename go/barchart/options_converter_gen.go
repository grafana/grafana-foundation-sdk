// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package barchart

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
)

// OptionsConverter accepts a `Options` object and generates the Go code to build this object using builders.
func OptionsConverter(input Options) string {
	calls := []string{
		`barchart.NewOptionsBuilder()`,
	}
	var buffer strings.Builder
	if input.XField != nil && *input.XField != "" {

		buffer.WriteString(`XField(`)
		arg0 := fmt.Sprintf("%#v", *input.XField)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.ColorByField != nil && *input.ColorByField != "" {

		buffer.WriteString(`ColorByField(`)
		arg0 := fmt.Sprintf("%#v", *input.ColorByField)
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

	if input.BarRadius != nil && *input.BarRadius != 0 {

		buffer.WriteString(`BarRadius(`)
		arg0 := fmt.Sprintf("%#v", *input.BarRadius)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.XTickLabelRotation != 0 {

		buffer.WriteString(`XTickLabelRotation(`)
		arg0 := fmt.Sprintf("%#v", input.XTickLabelRotation)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	{
		buffer.WriteString(`XTickLabelMaxLength(`)
		arg0 := fmt.Sprintf("%#v", input.XTickLabelMaxLength)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.XTickLabelSpacing != nil && *input.XTickLabelSpacing != 0 {

		buffer.WriteString(`XTickLabelSpacing(`)
		arg0 := fmt.Sprintf("%#v", *input.XTickLabelSpacing)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	{
		buffer.WriteString(`Stacking(`)
		arg0 := cog.Dump(input.Stacking)
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

	if input.BarWidth != 0.97 {

		buffer.WriteString(`BarWidth(`)
		arg0 := fmt.Sprintf("%#v", input.BarWidth)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.GroupWidth != 0.7 {

		buffer.WriteString(`GroupWidth(`)
		arg0 := fmt.Sprintf("%#v", input.GroupWidth)
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

	if input.Text != nil {

		buffer.WriteString(`Text(`)
		arg0 := common.VizTextDisplayOptionsConverter(*input.Text)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.FullHighlight != false {

		buffer.WriteString(`FullHighlight(`)
		arg0 := fmt.Sprintf("%#v", input.FullHighlight)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
