// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// MetricAggregationTypeConverter accepts a `MetricAggregationType` object and generates the Go code to build this object using builders.
func MetricAggregationTypeConverter(input MetricAggregationType) string {
	calls := []string{
		`elasticsearch.NewMetricAggregationTypeBuilder()`,
	}
	var buffer strings.Builder
	if input.PipelineMetricAggregationType != nil {

		buffer.WriteString(`PipelineMetricAggregationType(`)
		arg0 := cog.Dump(*input.PipelineMetricAggregationType)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
