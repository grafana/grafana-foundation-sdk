// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// ElasticsearchSumSettingsConverter accepts a `ElasticsearchSumSettings` object and generates the Go code to build this object using builders.
func ElasticsearchSumSettingsConverter(input ElasticsearchSumSettings) string {
	calls := []string{
		`elasticsearch.NewElasticsearchSumSettingsBuilder()`,
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
