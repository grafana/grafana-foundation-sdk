// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alertingv1beta1

import (
	"strings"
)

// RuleSequenceConverter accepts a `RuleSequence` object and generates the Go code to build this object using builders.
func RuleSequenceConverter(input RuleSequenceSpec) string {
	calls := []string{
		`alertingv1beta1.NewRuleSequenceBuilder()`,
	}
	var buffer strings.Builder

	{
		buffer.WriteString(`Trigger(`)
		arg0 := IntervalTriggerConverter(input.Trigger)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.RecordingRules != nil && len(input.RecordingRules) >= 1 {

		buffer.WriteString(`RecordingRules(`)
		tmparg0 := []string{}
		for _, arg1 := range input.RecordingRules {
			tmprecordingRulesarg1 := RuleRefConverter(arg1)
			tmparg0 = append(tmparg0, tmprecordingRulesarg1)
		}
		arg0 := "[]cog.Builder[alertingv1beta1.RuleRef]{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.AlertRules != nil && len(input.AlertRules) >= 1 {

		buffer.WriteString(`AlertRules(`)
		tmparg0 := []string{}
		for _, arg1 := range input.AlertRules {
			tmpalertRulesarg1 := RuleRefConverter(arg1)
			tmparg0 = append(tmparg0, tmpalertRulesarg1)
		}
		arg0 := "[]cog.Builder[alertingv1beta1.RuleRef]{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
