// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// ElasticsearchInlineScriptConverter accepts a `ElasticsearchInlineScript` object and generates the Go code to build this object using builders.
func ElasticsearchInlineScriptConverter(input ElasticsearchInlineScript) string {
	calls := []string{
		`elasticsearch.NewElasticsearchInlineScriptBuilder()`,
	}
	var buffer strings.Builder
	if input.Inline != nil && *input.Inline != "" {

		buffer.WriteString(`Inline(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Inline))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
