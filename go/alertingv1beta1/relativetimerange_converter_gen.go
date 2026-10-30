// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alertingv1beta1

import (
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// RelativeTimeRangeConverter accepts a `RelativeTimeRange` object and generates the Go code to build this object using builders.
func RelativeTimeRangeConverter(input RelativeTimeRange) string {
	calls := []string{
		`alertingv1beta1.NewRelativeTimeRangeBuilder()`,
	}
	var buffer strings.Builder

	{
		buffer.WriteString(`From(`)
		arg0 := cog.Dump(input.From)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	{
		buffer.WriteString(`To(`)
		arg0 := cog.Dump(input.To)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	return strings.Join(calls, ".\t\n")
}
