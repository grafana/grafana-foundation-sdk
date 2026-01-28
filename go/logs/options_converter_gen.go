// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package logs

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// OptionsConverter accepts a `Options` object and generates the Go code to build this object using builders.
func OptionsConverter(input Options) string {
	calls := []string{
		`logs.NewOptionsBuilder()`,
	}
	var buffer strings.Builder

	{
		buffer.WriteString(`ShowLabels(`)
		arg0 := fmt.Sprintf("%#v", input.ShowLabels)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	{
		buffer.WriteString(`ShowCommonLabels(`)
		arg0 := fmt.Sprintf("%#v", input.ShowCommonLabels)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	{
		buffer.WriteString(`ShowTime(`)
		arg0 := fmt.Sprintf("%#v", input.ShowTime)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	{
		buffer.WriteString(`ShowLogContextToggle(`)
		arg0 := fmt.Sprintf("%#v", input.ShowLogContextToggle)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	{
		buffer.WriteString(`WrapLogMessage(`)
		arg0 := fmt.Sprintf("%#v", input.WrapLogMessage)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	{
		buffer.WriteString(`PrettifyLogMessage(`)
		arg0 := fmt.Sprintf("%#v", input.PrettifyLogMessage)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	{
		buffer.WriteString(`EnableLogDetails(`)
		arg0 := fmt.Sprintf("%#v", input.EnableLogDetails)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	{
		buffer.WriteString(`SortOrder(`)
		arg0 := cog.Dump(input.SortOrder)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	{
		buffer.WriteString(`DedupStrategy(`)
		arg0 := cog.Dump(input.DedupStrategy)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.EnableInfiniteScrolling != nil {

		buffer.WriteString(`EnableInfiniteScrolling(`)
		arg0 := fmt.Sprintf("%#v", *input.EnableInfiniteScrolling)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.OnClickFilterLabel != nil {

		buffer.WriteString(`OnClickFilterLabel(`)
		arg0 := cog.Dump(input.OnClickFilterLabel)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.OnClickFilterOutLabel != nil {

		buffer.WriteString(`OnClickFilterOutLabel(`)
		arg0 := cog.Dump(input.OnClickFilterOutLabel)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.IsFilterLabelActive != nil {

		buffer.WriteString(`IsFilterLabelActive(`)
		arg0 := cog.Dump(input.IsFilterLabelActive)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.OnClickFilterString != nil {

		buffer.WriteString(`OnClickFilterString(`)
		arg0 := cog.Dump(input.OnClickFilterString)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.OnClickFilterOutString != nil {

		buffer.WriteString(`OnClickFilterOutString(`)
		arg0 := cog.Dump(input.OnClickFilterOutString)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.OnClickShowField != nil {

		buffer.WriteString(`OnClickShowField(`)
		arg0 := cog.Dump(input.OnClickShowField)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.OnClickHideField != nil {

		buffer.WriteString(`OnClickHideField(`)
		arg0 := cog.Dump(input.OnClickHideField)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.LogRowMenuIconsBefore != nil {

		buffer.WriteString(`LogRowMenuIconsBefore(`)
		arg0 := cog.Dump(input.LogRowMenuIconsBefore)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.LogRowMenuIconsAfter != nil {

		buffer.WriteString(`LogRowMenuIconsAfter(`)
		arg0 := cog.Dump(input.LogRowMenuIconsAfter)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.OnNewLogsReceived != nil {

		buffer.WriteString(`OnNewLogsReceived(`)
		arg0 := cog.Dump(input.OnNewLogsReceived)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.DisplayedFields != nil && len(input.DisplayedFields) >= 1 {

		buffer.WriteString(`DisplayedFields(`)
		tmparg0 := []string{}
		for _, arg1 := range input.DisplayedFields {
			tmpdisplayedFieldsarg1 := fmt.Sprintf("%#v", arg1)
			tmparg0 = append(tmparg0, tmpdisplayedFieldsarg1)
		}
		arg0 := "[]string{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
