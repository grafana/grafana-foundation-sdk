// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	"fmt"
	"strings"
)

// MetricNamespaceQueryConverter accepts a `MetricNamespaceQuery` object and generates the Go code to build this object using builders.
func MetricNamespaceQueryConverter(input MetricNamespaceQuery) string {
	calls := []string{
		`azuremonitor.NewMetricNamespaceQueryBuilder()`,
	}
	var buffer strings.Builder
	if input.RawQuery != nil && *input.RawQuery != "" {

		buffer.WriteString(`RawQuery(`)
		arg0 := fmt.Sprintf("%#v", *input.RawQuery)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Subscription != "" {

		buffer.WriteString(`Subscription(`)
		arg0 := fmt.Sprintf("%#v", input.Subscription)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.ResourceGroup != "" {

		buffer.WriteString(`ResourceGroup(`)
		arg0 := fmt.Sprintf("%#v", input.ResourceGroup)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.MetricNamespace != nil && *input.MetricNamespace != "" {

		buffer.WriteString(`MetricNamespace(`)
		arg0 := fmt.Sprintf("%#v", *input.MetricNamespace)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.ResourceName != nil && *input.ResourceName != "" {

		buffer.WriteString(`ResourceName(`)
		arg0 := fmt.Sprintf("%#v", *input.ResourceName)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
