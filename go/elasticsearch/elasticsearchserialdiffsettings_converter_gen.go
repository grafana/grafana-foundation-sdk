// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// ElasticsearchSerialDiffSettingsConverter accepts a `ElasticsearchSerialDiffSettings` object and generates the Go code to build this object using builders.
func ElasticsearchSerialDiffSettingsConverter(input ElasticsearchSerialDiffSettings) string {
	calls := []string{
		`elasticsearch.NewElasticsearchSerialDiffSettingsBuilder()`,
	}
	var buffer strings.Builder
	if input.Lag != nil && *input.Lag != "" {

		buffer.WriteString(`Lag(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Lag))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
