// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alertingv1beta1

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// AlertRuleStatusConverter accepts a `AlertRuleStatus` object and generates the Go code to build this object using builders.
func AlertRuleStatusConverter(input AlertRuleStatus) string {
	calls := []string{
		`alertingv1beta1.NewAlertRuleStatusBuilder()`,
	}
	var buffer strings.Builder
	if input.Health != nil {

		buffer.WriteString(`Health(`)
		arg0 := cog.Dump(*input.Health)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.State != nil {

		buffer.WriteString(`State(`)
		arg0 := cog.Dump(*input.State)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.StateReason != nil {

		buffer.WriteString(`StateReason(`)
		arg0 := cog.Dump(*input.StateReason)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.LastEvaluationTime != nil {

		buffer.WriteString(`LastEvaluationTime(`)
		arg0 := fmt.Sprintf("%#v", *input.LastEvaluationTime)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.EvaluationDuration != nil {

		buffer.WriteString(`EvaluationDuration(`)
		arg0 := fmt.Sprintf("%#v", *input.EvaluationDuration)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.LastError != nil && *input.LastError != "" {

		buffer.WriteString(`LastError(`)
		arg0 := fmt.Sprintf("%#v", *input.LastError)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Totals != nil {

		buffer.WriteString(`Totals(`)
		arg0 := AlertRuleInstanceTotalsConverter(*input.Totals)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
