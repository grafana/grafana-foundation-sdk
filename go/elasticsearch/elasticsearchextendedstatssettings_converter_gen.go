// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// ElasticsearchExtendedStatsSettingsConverter accepts a `ElasticsearchExtendedStatsSettings` object and generates the Go code to build this object using builders.
func ElasticsearchExtendedStatsSettingsConverter(input ElasticsearchExtendedStatsSettings) string {
	calls := []string{
		`elasticsearch.NewElasticsearchExtendedStatsSettingsBuilder()`,
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
	if input.Sigma != nil && *input.Sigma != "" {

		buffer.WriteString(`Sigma(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Sigma))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
