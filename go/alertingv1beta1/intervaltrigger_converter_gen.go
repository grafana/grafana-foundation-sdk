// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alertingv1beta1

import (
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// IntervalTriggerConverter accepts a `IntervalTrigger` object and generates the Go code to build this object using builders.
func IntervalTriggerConverter(input IntervalTrigger) string {
	calls := []string{
		`alertingv1beta1.NewIntervalTriggerBuilder()`,
	}
	var buffer strings.Builder

	{
		buffer.WriteString(`Interval(`)
		arg0 := cog.Dump(input.Interval)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	return strings.Join(calls, ".\t\n")
}
