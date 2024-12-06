// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alerting

import (
	"fmt"
	"strings"
)

// RecordRuleConverter accepts a `RecordRule` object and generates the Go code to build this object using builders.
func RecordRuleConverter(input RecordRule) string {
	calls := []string{
		`alerting.NewRecordRuleBuilder()`,
	}
	var buffer strings.Builder
	if input.From != "" {

		buffer.WriteString(`From(`)
		arg0 := fmt.Sprintf("%#v", input.From)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Metric != "" {

		buffer.WriteString(`Metric(`)
		arg0 := fmt.Sprintf("%#v", input.Metric)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
