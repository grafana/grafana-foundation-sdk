// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alerting

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

func MuteTimingConverter(input MuteTiming) string {
	calls := []string{
		`alerting.NewMuteTimingBuilder()`,
	}
	var buffer strings.Builder
	if input.Name != nil && *input.Name != "" {

		buffer.WriteString(`Name(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Name))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.TimeIntervals != nil && len(input.TimeIntervals) >= 1 {

		buffer.WriteString(`TimeIntervals(`)
		tmparg0 := []string{}
		for _, arg1 := range input.TimeIntervals {
			tmptime_intervalsarg1 := TimeIntervalConverter(arg1)
			tmparg0 = append(tmparg0, tmptime_intervalsarg1)
		}
		arg0 := "[]cog.Builder[alerting.TimeInterval]{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}