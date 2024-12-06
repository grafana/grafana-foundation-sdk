// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// MetricDefinitionsQueryConverter accepts a `MetricDefinitionsQuery` object and generates the Go code to build this object using builders.
func MetricDefinitionsQueryConverter(input MetricDefinitionsQuery) string {
	calls := []string{
		`azuremonitor.NewMetricDefinitionsQueryBuilder()`,
	}
	var buffer strings.Builder
	if input.RawQuery != nil && *input.RawQuery != "" {

		buffer.WriteString(`RawQuery(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.RawQuery))
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
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.MetricNamespace))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.ResourceName != nil && *input.ResourceName != "" {

		buffer.WriteString(`ResourceName(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.ResourceName))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
