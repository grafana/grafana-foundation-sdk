// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// InlineScriptConverter accepts a `InlineScript` object and generates the Go code to build this object using builders.
func InlineScriptConverter(input InlineScript) string {
	calls := []string{
		`elasticsearch.NewInlineScriptBuilder()`,
	}
	var buffer strings.Builder
	if input.String != nil && *input.String != "" {

		buffer.WriteString(`String(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.String))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.ElasticsearchInlineScript != nil {

		buffer.WriteString(`ElasticsearchInlineScript(`)
		arg0 := ElasticsearchInlineScriptConverter(*input.ElasticsearchInlineScript)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
