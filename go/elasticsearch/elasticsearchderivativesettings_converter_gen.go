// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// ElasticsearchDerivativeSettingsConverter accepts a `ElasticsearchDerivativeSettings` object and generates the Go code to build this object using builders.
func ElasticsearchDerivativeSettingsConverter(input ElasticsearchDerivativeSettings) string {
	calls := []string{
		`elasticsearch.NewElasticsearchDerivativeSettingsBuilder()`,
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

	return strings.Join(calls, ".\t\n")
}
