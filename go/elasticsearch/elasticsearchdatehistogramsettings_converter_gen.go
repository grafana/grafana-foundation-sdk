// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// ElasticsearchDateHistogramSettingsConverter accepts a `ElasticsearchDateHistogramSettings` object and generates the Go code to build this object using builders.
func ElasticsearchDateHistogramSettingsConverter(input ElasticsearchDateHistogramSettings) string {
	calls := []string{
		`elasticsearch.NewElasticsearchDateHistogramSettingsBuilder()`,
	}
	var buffer strings.Builder
	if input.Interval != nil && *input.Interval != "" {

		buffer.WriteString(`Interval(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Interval))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.MinDocCount != nil && *input.MinDocCount != "" {

		buffer.WriteString(`MinDocCount(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.MinDocCount))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.TrimEdges != nil && *input.TrimEdges != "" {

		buffer.WriteString(`TrimEdges(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.TrimEdges))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Offset != nil && *input.Offset != "" {

		buffer.WriteString(`Offset(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Offset))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.TimeZone != nil && *input.TimeZone != "" {

		buffer.WriteString(`TimeZone(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.TimeZone))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
