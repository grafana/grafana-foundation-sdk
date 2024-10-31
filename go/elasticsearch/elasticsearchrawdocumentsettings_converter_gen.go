// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// ElasticsearchRawDocumentSettingsConverter accepts a `ElasticsearchRawDocumentSettings` object and generates the Go code to build this object using builders.
func ElasticsearchRawDocumentSettingsConverter(input ElasticsearchRawDocumentSettings) string {
	calls := []string{
		`elasticsearch.NewElasticsearchRawDocumentSettingsBuilder()`,
	}
	var buffer strings.Builder
	if input.Size != nil && *input.Size != "" {

		buffer.WriteString(`Size(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Size))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
