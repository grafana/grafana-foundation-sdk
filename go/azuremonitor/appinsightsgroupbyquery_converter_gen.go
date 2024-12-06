// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// AppInsightsGroupByQueryConverter accepts a `AppInsightsGroupByQuery` object and generates the Go code to build this object using builders.
func AppInsightsGroupByQueryConverter(input AppInsightsGroupByQuery) string {
	calls := []string{
		`azuremonitor.NewAppInsightsGroupByQueryBuilder()`,
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
	if input.MetricName != "" {

		buffer.WriteString(`MetricName(`)
		arg0 := fmt.Sprintf("%#v", input.MetricName)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
