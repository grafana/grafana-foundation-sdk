// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alerting

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// TimeRangeConverter accepts a `TimeRange` object and generates the Go code to build this object using builders.
func TimeRangeConverter(input TimeRange) string {
	calls := []string{
		`alerting.NewTimeRangeBuilder()`,
	}
	var buffer strings.Builder
	if input.From != nil {

		buffer.WriteString(`From(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.From))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.To != nil {

		buffer.WriteString(`To(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.To))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
