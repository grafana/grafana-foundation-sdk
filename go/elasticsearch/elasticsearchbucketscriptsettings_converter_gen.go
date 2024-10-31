// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	"strings"
)

// ElasticsearchBucketScriptSettingsConverter accepts a `ElasticsearchBucketScriptSettings` object and generates the Go code to build this object using builders.
func ElasticsearchBucketScriptSettingsConverter(input ElasticsearchBucketScriptSettings) string {
	calls := []string{
		`elasticsearch.NewElasticsearchBucketScriptSettingsBuilder()`,
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

	return strings.Join(calls, ".\t\n")
}
