// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package statushistory

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
)

// OptionsConverter accepts a `Options` object and generates the Go code to build this object using builders.
func OptionsConverter(input Options) string {
	calls := []string{
		`statushistory.NewOptionsBuilder()`,
	}
	var buffer strings.Builder
	if input.RowHeight != 0.9 {

		buffer.WriteString(`RowHeight(`)
		arg0 := fmt.Sprintf("%#v", input.RowHeight)
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

	if input.ColWidth != nil && *input.ColWidth != 0.9 {

		buffer.WriteString(`ColWidth(`)
		arg0 := fmt.Sprintf("%#v", *input.ColWidth)
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

	if input.Timezone != nil && len(input.Timezone) >= 1 {

		buffer.WriteString(`Timezone(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Timezone {
			tmptimezonearg1 := cog.Dump(arg1)
			tmparg0 = append(tmparg0, tmptimezonearg1)
		}
		arg0 := "[]common.TimeZone{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.PerPage != nil && *input.PerPage != 20 {

		buffer.WriteString(`PerPage(`)
		arg0 := fmt.Sprintf("%#v", *input.PerPage)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
