// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alertingv1beta1

import (
	"fmt"
	"strings"
)

// AlertRuleInstanceTotalsConverter accepts a `AlertRuleInstanceTotals` object and generates the Go code to build this object using builders.
func AlertRuleInstanceTotalsConverter(input AlertRuleInstanceTotals) string {
	calls := []string{
		`alertingv1beta1.NewAlertRuleInstanceTotalsBuilder()`,
	}
	var buffer strings.Builder
	if input.Healthy != nil {

		buffer.WriteString(`Healthy(`)
		arg0 := fmt.Sprintf("%#v", *input.Healthy)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Firing != nil {

		buffer.WriteString(`Firing(`)
		arg0 := fmt.Sprintf("%#v", *input.Firing)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Pending != nil {

		buffer.WriteString(`Pending(`)
		arg0 := fmt.Sprintf("%#v", *input.Pending)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Recovering != nil {

		buffer.WriteString(`Recovering(`)
		arg0 := fmt.Sprintf("%#v", *input.Recovering)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Nodata != nil {

		buffer.WriteString(`Nodata(`)
		arg0 := fmt.Sprintf("%#v", *input.Nodata)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Error != nil {

		buffer.WriteString(`Error(`)
		arg0 := fmt.Sprintf("%#v", *input.Error)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
