// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// ElasticsearchMovingAverageHoltWintersModelSettingsSettingsConverter accepts a `ElasticsearchMovingAverageHoltWintersModelSettingsSettings` object and generates the Go code to build this object using builders.
func ElasticsearchMovingAverageHoltWintersModelSettingsSettingsConverter(input ElasticsearchMovingAverageHoltWintersModelSettingsSettings) string {
	calls := []string{
		`elasticsearch.NewElasticsearchMovingAverageHoltWintersModelSettingsSettingsBuilder()`,
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
	if input.Beta != nil && *input.Beta != "" {

		buffer.WriteString(`Beta(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Beta))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Gamma != nil && *input.Gamma != "" {

		buffer.WriteString(`Gamma(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Gamma))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Period != nil && *input.Period != "" {

		buffer.WriteString(`Period(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Period))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Pad != nil {

		buffer.WriteString(`Pad(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Pad))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
