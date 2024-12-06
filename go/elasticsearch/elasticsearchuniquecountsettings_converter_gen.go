// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// ElasticsearchUniqueCountSettingsConverter accepts a `ElasticsearchUniqueCountSettings` object and generates the Go code to build this object using builders.
func ElasticsearchUniqueCountSettingsConverter(input ElasticsearchUniqueCountSettings) string {
	calls := []string{
		`elasticsearch.NewElasticsearchUniqueCountSettingsBuilder()`,
	}
	var buffer strings.Builder
	if input.PrecisionThreshold != nil && *input.PrecisionThreshold != "" {

		buffer.WriteString(`PrecisionThreshold(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.PrecisionThreshold))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
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
