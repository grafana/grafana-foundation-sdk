// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// ElasticsearchMovingAverageEWMAModelSettingsSettingsConverter accepts a `ElasticsearchMovingAverageEWMAModelSettingsSettings` object and generates the Go code to build this object using builders.
func ElasticsearchMovingAverageEWMAModelSettingsSettingsConverter(input ElasticsearchMovingAverageEWMAModelSettingsSettings) string {
	calls := []string{
		`elasticsearch.NewElasticsearchMovingAverageEWMAModelSettingsSettingsBuilder()`,
	}
	var buffer strings.Builder
	if input.Alpha != nil && *input.Alpha != "" {

		buffer.WriteString(`Alpha(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Alpha))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
