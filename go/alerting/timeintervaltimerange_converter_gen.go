// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alerting

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// TimeIntervalTimeRangeConverter accepts a `TimeIntervalTimeRange` object and generates the Go code to build this object using builders.
func TimeIntervalTimeRangeConverter(input TimeIntervalTimeRange) string {
	calls := []string{
		`alerting.NewTimeIntervalTimeRangeBuilder()`,
	}
	var buffer strings.Builder
	if input.EndTime != nil && *input.EndTime != "" {

		buffer.WriteString(`EndTime(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.EndTime))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.StartTime != nil && *input.StartTime != "" {

		buffer.WriteString(`StartTime(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.StartTime))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
