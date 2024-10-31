// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// ElasticsearchMovingFunctionSettingsConverter accepts a `ElasticsearchMovingFunctionSettings` object and generates the Go code to build this object using builders.
func ElasticsearchMovingFunctionSettingsConverter(input ElasticsearchMovingFunctionSettings) string {
	calls := []string{
		`elasticsearch.NewElasticsearchMovingFunctionSettingsBuilder()`,
	}
	var buffer strings.Builder
	if input.Window != nil && *input.Window != "" {

		buffer.WriteString(`Window(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Window))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Script != nil {

		buffer.WriteString(`Script(`)
		arg0 := InlineScriptConverter(*input.Script)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Shift != nil && *input.Shift != "" {

		buffer.WriteString(`Shift(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Shift))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
