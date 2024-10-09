// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

func ElasticsearchMetricAggregationWithMissingSupportSettingsConverter(input ElasticsearchMetricAggregationWithMissingSupportSettings) string {
	calls := []string{
		`elasticsearch.NewElasticsearchMetricAggregationWithMissingSupportSettingsBuilder()`,
	}
	var buffer strings.Builder
	if input.Missing != nil && *input.Missing != "" {

		buffer.WriteString(`Missing(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Missing))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
