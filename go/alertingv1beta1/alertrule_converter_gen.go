// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alertingv1beta1

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// AlertRuleConverter accepts a `AlertRule` object and generates the Go code to build this object using builders.
func AlertRuleConverter(input AlertRuleSpec) string {
	calls := []string{
		`alertingv1beta1.NewAlertRuleBuilder()`,
	}
	var buffer strings.Builder
	if input.Title != "" {

		buffer.WriteString(`Title(`)
		arg0 := fmt.Sprintf("%#v", input.Title)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Paused != nil {

		buffer.WriteString(`Paused(`)
		arg0 := fmt.Sprintf("%#v", *input.Paused)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	{
		buffer.WriteString(`Trigger(`)
		arg0 := IntervalTriggerConverter(input.Trigger)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.Labels != nil {

		buffer.WriteString(`Labels(`)
		arg0 := "map[string]alertingv1beta1.TemplateString{"
		for key, arg1 := range input.Labels {
			tmplabelsarg1 := cog.Dump(arg1)
			arg0 += "\t" + fmt.Sprintf("%#v", key) + ": " + tmplabelsarg1 + ","
		}
		arg0 += "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Annotations != nil {

		buffer.WriteString(`Annotations(`)
		arg0 := "map[string]alertingv1beta1.TemplateString{"
		for key, arg1 := range input.Annotations {
			tmpannotationsarg1 := cog.Dump(arg1)
			arg0 += "\t" + fmt.Sprintf("%#v", key) + ": " + tmpannotationsarg1 + ","
		}
		arg0 += "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.For != nil && *input.For != "" {

		buffer.WriteString(`For(`)
		arg0 := fmt.Sprintf("%#v", *input.For)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.KeepFiringFor != nil && *input.KeepFiringFor != "" {

		buffer.WriteString(`KeepFiringFor(`)
		arg0 := fmt.Sprintf("%#v", *input.KeepFiringFor)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.MissingSeriesEvalsToResolve != nil {

		buffer.WriteString(`MissingSeriesEvalsToResolve(`)
		arg0 := fmt.Sprintf("%#v", *input.MissingSeriesEvalsToResolve)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	{
		buffer.WriteString(`NoDataState(`)
		arg0 := cog.Dump(input.NoDataState)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	{
		buffer.WriteString(`ExecErrState(`)
		arg0 := cog.Dump(input.ExecErrState)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.NotificationSettings != nil {

		buffer.WriteString(`NotificationSettings(`)
		arg0 := cog.Dump(*input.NotificationSettings)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	{
		buffer.WriteString(`Expressions(`)
		arg0 := cog.Dump(input.Expressions)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.PanelRef != nil {

		buffer.WriteString(`PanelRef(`)
		arg0 := PanelRefConverter(*input.PanelRef)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
