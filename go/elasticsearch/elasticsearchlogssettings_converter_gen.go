// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// ElasticsearchLogsSettingsConverter accepts a `ElasticsearchLogsSettings` object and generates the Go code to build this object using builders.
func ElasticsearchLogsSettingsConverter(input ElasticsearchLogsSettings) string {
	calls := []string{
		`elasticsearch.NewElasticsearchLogsSettingsBuilder()`,
	}
	var buffer strings.Builder
	if input.Limit != nil && *input.Limit != "" {

		buffer.WriteString(`Limit(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Limit))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
