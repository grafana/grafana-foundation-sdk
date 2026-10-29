// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alertingv0alpha1

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// SimplifiedRoutingConverter accepts a `SimplifiedRouting` object and generates the Go code to build this object using builders.
func SimplifiedRoutingConverter(input SimplifiedRouting) string {
	calls := []string{
		`alertingv0alpha1.NewSimplifiedRoutingBuilder()`,
	}
	var buffer strings.Builder
	if input.Receiver != "" {

		buffer.WriteString(`Receiver(`)
		arg0 := fmt.Sprintf("%#v", input.Receiver)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.GroupBy != nil && len(input.GroupBy) >= 1 {

		buffer.WriteString(`GroupBy(`)
		tmparg0 := []string{}
		for _, arg1 := range input.GroupBy {
			tmpgroupByarg1 := fmt.Sprintf("%#v", arg1)
			tmparg0 = append(tmparg0, tmpgroupByarg1)
		}
		arg0 := "[]string{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.GroupWait != nil {

		buffer.WriteString(`GroupWait(`)
		arg0 := cog.Dump(*input.GroupWait)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.GroupInterval != nil {

		buffer.WriteString(`GroupInterval(`)
		arg0 := cog.Dump(*input.GroupInterval)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.RepeatInterval != nil {

		buffer.WriteString(`RepeatInterval(`)
		arg0 := cog.Dump(*input.RepeatInterval)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.MuteTimeIntervals != nil && len(input.MuteTimeIntervals) >= 1 {

		buffer.WriteString(`MuteTimeIntervals(`)
		tmparg0 := []string{}
		for _, arg1 := range input.MuteTimeIntervals {
			tmpmuteTimeIntervalsarg1 := cog.Dump(arg1)
			tmparg0 = append(tmparg0, tmpmuteTimeIntervalsarg1)
		}
		arg0 := "[]alertingv0alpha1.TimeIntervalRef{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.ActiveTimeIntervals != nil && len(input.ActiveTimeIntervals) >= 1 {

		buffer.WriteString(`ActiveTimeIntervals(`)
		tmparg0 := []string{}
		for _, arg1 := range input.ActiveTimeIntervals {
			tmpactiveTimeIntervalsarg1 := cog.Dump(arg1)
			tmparg0 = append(tmparg0, tmpactiveTimeIntervalsarg1)
		}
		arg0 := "[]alertingv0alpha1.TimeIntervalRef{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
