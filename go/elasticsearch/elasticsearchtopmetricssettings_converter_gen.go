// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// ElasticsearchTopMetricsSettingsConverter accepts a `ElasticsearchTopMetricsSettings` object and generates the Go code to build this object using builders.
func ElasticsearchTopMetricsSettingsConverter(input ElasticsearchTopMetricsSettings) string {
	calls := []string{
		`elasticsearch.NewElasticsearchTopMetricsSettingsBuilder()`,
	}
	var buffer strings.Builder
	if input.Order != nil && *input.Order != "" {

		buffer.WriteString(`Order(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Order))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.OrderBy != nil && *input.OrderBy != "" {

		buffer.WriteString(`OrderBy(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.OrderBy))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Metrics != nil && len(input.Metrics) >= 1 {

		buffer.WriteString(`Metrics(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Metrics {
			tmpmetricsarg1 := fmt.Sprintf("%#v", arg1)
			tmparg0 = append(tmparg0, tmpmetricsarg1)
		}
		arg0 := "[]string{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
