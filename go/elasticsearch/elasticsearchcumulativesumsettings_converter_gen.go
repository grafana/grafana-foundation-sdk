// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// ElasticsearchCumulativeSumSettingsConverter accepts a `ElasticsearchCumulativeSumSettings` object and generates the Go code to build this object using builders.
func ElasticsearchCumulativeSumSettingsConverter(input ElasticsearchCumulativeSumSettings) string {
	calls := []string{
		`elasticsearch.NewElasticsearchCumulativeSumSettingsBuilder()`,
	}
	var buffer strings.Builder
	if input.Format != nil && *input.Format != "" {

		buffer.WriteString(`Format(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Format))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
