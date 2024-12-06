// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alerting

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// NotificationPolicyConverter accepts a `NotificationPolicy` object and generates the Go code to build this object using builders.
func NotificationPolicyConverter(input NotificationPolicy) string {
	calls := []string{
		`alerting.NewNotificationPolicyBuilder()`,
	}
	var buffer strings.Builder
	if input.ActiveTimeIntervals != nil && len(input.ActiveTimeIntervals) >= 1 {

		buffer.WriteString(`ActiveTimeIntervals(`)
		tmparg0 := []string{}
		for _, arg1 := range input.ActiveTimeIntervals {
			tmpactive_time_intervalsarg1 := fmt.Sprintf("%#v", arg1)
			tmparg0 = append(tmparg0, tmpactive_time_intervalsarg1)
		}
		arg0 := "[]string{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Continue != nil {

		buffer.WriteString(`Continue(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Continue))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
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
	if input.Match != nil {

		buffer.WriteString(`Match(`)
		arg0 := "map[string]string{"
		for key, arg1 := range input.Match {
			tmpmatcharg1 := fmt.Sprintf("%#v", arg1)
			arg0 += "\t" + fmt.Sprintf("%#v", key) + ": " + tmpmatcharg1 + ","
		}
		arg0 += "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.MatchRe != nil {

		buffer.WriteString(`MatchRe(`)
		arg0 := cog.Dump(*input.MatchRe)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Matchers != nil {

		buffer.WriteString(`Matchers(`)
		arg0 := cog.Dump(*input.Matchers)
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
	if input.ObjectMatchers != nil {

		buffer.WriteString(`ObjectMatchers(`)
		arg0 := cog.Dump(*input.ObjectMatchers)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Provenance != nil {

		buffer.WriteString(`Provenance(`)
		arg0 := cog.Dump(*input.Provenance)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Receiver != nil && *input.Receiver != "" {

		buffer.WriteString(`Receiver(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Receiver))
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
	if input.Routes != nil && len(input.Routes) >= 1 {

		buffer.WriteString(`Routes(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Routes {
			tmproutesarg1 := NotificationPolicyConverter(arg1)
			tmparg0 = append(tmparg0, tmproutesarg1)
		}
		arg0 := "[]cog.Builder[alerting.NotificationPolicy]{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
