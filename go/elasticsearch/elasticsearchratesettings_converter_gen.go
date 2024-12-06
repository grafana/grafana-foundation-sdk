// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// ElasticsearchRateSettingsConverter accepts a `ElasticsearchRateSettings` object and generates the Go code to build this object using builders.
func ElasticsearchRateSettingsConverter(input ElasticsearchRateSettings) string {
	calls := []string{
		`elasticsearch.NewElasticsearchRateSettingsBuilder()`,
	}
	var buffer strings.Builder
	if input.Unit != nil && *input.Unit != "" {

		buffer.WriteString(`Unit(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Unit))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Mode != nil && *input.Mode != "" {

		buffer.WriteString(`Mode(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Mode))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
