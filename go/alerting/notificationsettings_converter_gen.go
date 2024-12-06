// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alerting

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// NotificationSettingsConverter accepts a `NotificationSettings` object and generates the Go code to build this object using builders.
func NotificationSettingsConverter(input NotificationSettings) string {
	calls := []string{
		`alerting.NewNotificationSettingsBuilder()`,
	}
	var buffer strings.Builder
	if input.GroupBy != nil && len(input.GroupBy) >= 1 {

		buffer.WriteString(`GroupBy(`)
		tmparg0 := []string{}
		for _, arg1 := range input.GroupBy {
			tmpgroup_byarg1 := fmt.Sprintf("%#v", arg1)
			tmparg0 = append(tmparg0, tmpgroup_byarg1)
		}
		arg0 := "[]string{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.GroupInterval != nil && *input.GroupInterval != "" {

		buffer.WriteString(`GroupInterval(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.GroupInterval))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.GroupWait != nil && *input.GroupWait != "" {

		buffer.WriteString(`GroupWait(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.GroupWait))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.MuteTimeIntervals != nil && len(input.MuteTimeIntervals) >= 1 {

		buffer.WriteString(`MuteTimeIntervals(`)
		tmparg0 := []string{}
		for _, arg1 := range input.MuteTimeIntervals {
			tmpmute_time_intervalsarg1 := fmt.Sprintf("%#v", arg1)
			tmparg0 = append(tmparg0, tmpmute_time_intervalsarg1)
		}
		arg0 := "[]string{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Receiver != "" {

		buffer.WriteString(`Receiver(`)
		arg0 := fmt.Sprintf("%#v", input.Receiver)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.RepeatInterval != nil && *input.RepeatInterval != "" {

		buffer.WriteString(`RepeatInterval(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.RepeatInterval))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
