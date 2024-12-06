// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// HistogramSettingsConverter accepts a `HistogramSettings` object and generates the Go code to build this object using builders.
func HistogramSettingsConverter(input HistogramSettings) string {
	calls := []string{
		`elasticsearch.NewHistogramSettingsBuilder()`,
	}
	var buffer strings.Builder
	if input.Interval != nil && *input.Interval != "" {

		buffer.WriteString(`Interval(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Interval))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.MinDocCount != nil && *input.MinDocCount != "" {

		buffer.WriteString(`MinDocCount(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.MinDocCount))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
