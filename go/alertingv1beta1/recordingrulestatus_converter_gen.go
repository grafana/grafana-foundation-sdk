// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alertingv1beta1

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// RecordingRuleStatusConverter accepts a `RecordingRuleStatus` object and generates the Go code to build this object using builders.
func RecordingRuleStatusConverter(input RecordingRuleStatus) string {
	calls := []string{
		`alertingv1beta1.NewRecordingRuleStatusBuilder()`,
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

	return strings.Join(calls, ".\t\n")
}
