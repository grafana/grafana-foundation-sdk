// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package gauge

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
)

// OptionsConverter accepts a `Options` object and generates the Go code to build this object using builders.
func OptionsConverter(input Options) string {
	calls := []string{
		`gauge.NewOptionsBuilder()`,
	}
	var buffer strings.Builder
	if input.ShowThresholdLabels != false {

		buffer.WriteString(`ShowThresholdLabels(`)
		arg0 := fmt.Sprintf("%#v", input.ShowThresholdLabels)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.ShowThresholdMarkers != true {

		buffer.WriteString(`ShowThresholdMarkers(`)
		arg0 := fmt.Sprintf("%#v", input.ShowThresholdMarkers)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	{
		buffer.WriteString(`Sizing(`)
		arg0 := cog.Dump(input.Sizing)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.MinVizWidth != 75 {

		buffer.WriteString(`MinVizWidth(`)
		arg0 := fmt.Sprintf("%#v", input.MinVizWidth)
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
	if input.MinVizHeight != 75 {

		buffer.WriteString(`MinVizHeight(`)
		arg0 := fmt.Sprintf("%#v", input.MinVizHeight)
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
