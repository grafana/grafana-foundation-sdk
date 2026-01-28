// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package bargauge

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
)

// OptionsConverter accepts a `Options` object and generates the Go code to build this object using builders.
func OptionsConverter(input Options) string {
	calls := []string{
		`bargauge.NewOptionsBuilder()`,
	}
	var buffer strings.Builder

	{
		buffer.WriteString(`DisplayMode(`)
		arg0 := cog.Dump(input.DisplayMode)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	{
		buffer.WriteString(`ValueMode(`)
		arg0 := cog.Dump(input.ValueMode)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.ShowUnfilled != true {

		buffer.WriteString(`ShowUnfilled(`)
		arg0 := fmt.Sprintf("%#v", input.ShowUnfilled)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.MinVizWidth != 0 {

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
	if input.MinVizHeight != 10 {

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
