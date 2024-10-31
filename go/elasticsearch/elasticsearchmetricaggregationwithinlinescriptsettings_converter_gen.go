// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	"strings"
)

// ElasticsearchMetricAggregationWithInlineScriptSettingsConverter accepts a `ElasticsearchMetricAggregationWithInlineScriptSettings` object and generates the Go code to build this object using builders.
func ElasticsearchMetricAggregationWithInlineScriptSettingsConverter(input ElasticsearchMetricAggregationWithInlineScriptSettings) string {
	calls := []string{
		`elasticsearch.NewElasticsearchMetricAggregationWithInlineScriptSettingsBuilder()`,
	}
	var buffer strings.Builder
	if input.Script != nil {

		buffer.WriteString(`Script(`)
		arg0 := InlineScriptConverter(*input.Script)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
