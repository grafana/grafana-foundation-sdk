// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package googlecloudmonitoring

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// TimeSeriesQueryConverter accepts a `TimeSeriesQuery` object and generates the Go code to build this object using builders.
func TimeSeriesQueryConverter(input TimeSeriesQuery) string {
	calls := []string{
		`googlecloudmonitoring.NewTimeSeriesQueryBuilder()`,
	}
	var buffer strings.Builder
	if input.ProjectName != "" {

		buffer.WriteString(`ProjectName(`)
		arg0 := fmt.Sprintf("%#v", input.ProjectName)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Query != "" {

		buffer.WriteString(`Query(`)
		arg0 := fmt.Sprintf("%#v", input.Query)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.GraphPeriod != nil && *input.GraphPeriod != "" && *input.GraphPeriod != "disabled" {

		buffer.WriteString(`GraphPeriod(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.GraphPeriod))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
