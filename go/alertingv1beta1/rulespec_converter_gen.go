// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alertingv1beta1

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// RuleSpecConverter accepts a `RuleSpec` object and generates the Go code to build this object using builders.
func RuleSpecConverter(input RuleSpec) string {
	calls := []string{
		`alertingv1beta1.NewRuleSpecBuilder()`,
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

	{
		buffer.WriteString(`Expressions(`)
		arg0 := cog.Dump(input.Expressions)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	return strings.Join(calls, ".\t\n")
}
