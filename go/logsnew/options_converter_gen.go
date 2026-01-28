// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package logsnew

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// OptionsConverter accepts a `Options` object and generates the Go code to build this object using builders.
func OptionsConverter(input Options) string {
	calls := []string{
		`logsnew.NewOptionsBuilder()`,
	}
	var buffer strings.Builder

	{
		buffer.WriteString(`ShowTime(`)
		arg0 := fmt.Sprintf("%#v", input.ShowTime)
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
	if input.OnNewLogsReceived != nil {

		buffer.WriteString(`OnNewLogsReceived(`)
		arg0 := cog.Dump(input.OnNewLogsReceived)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
