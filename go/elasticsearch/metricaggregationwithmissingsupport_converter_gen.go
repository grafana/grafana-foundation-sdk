// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// MetricAggregationWithMissingSupportConverter accepts a `MetricAggregationWithMissingSupport` object and generates the Go code to build this object using builders.
func MetricAggregationWithMissingSupportConverter(input MetricAggregationWithMissingSupport) string {
	calls := []string{
		`elasticsearch.NewMetricAggregationWithMissingSupportBuilder()`,
	}
	var buffer strings.Builder
	if input.Settings != nil {

		buffer.WriteString(`Settings(`)
		arg0 := ElasticsearchMetricAggregationWithMissingSupportSettingsConverter(*input.Settings)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	{
		buffer.WriteString(`Type(`)
		arg0 := MetricAggregationTypeConverter(input.Type)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.Id != "" {

		buffer.WriteString(`Id(`)
		arg0 := fmt.Sprintf("%#v", input.Id)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Hide != nil {

		buffer.WriteString(`Hide(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Hide))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
