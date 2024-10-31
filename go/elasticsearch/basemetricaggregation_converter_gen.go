// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// BaseMetricAggregationConverter accepts a `BaseMetricAggregation` object and generates the Go code to build this object using builders.
func BaseMetricAggregationConverter(input BaseMetricAggregation) string {
	calls := []string{
		`elasticsearch.NewBaseMetricAggregationBuilder()`,
	}
	var buffer strings.Builder

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
