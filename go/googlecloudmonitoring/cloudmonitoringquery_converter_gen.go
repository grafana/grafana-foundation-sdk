// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package googlecloudmonitoring

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// CloudMonitoringQueryConverter accepts a `CloudMonitoringQuery` object and generates the Go code to build this object using builders.
func CloudMonitoringQueryConverter(input CloudMonitoringQuery) string {
	calls := []string{
		`googlecloudmonitoring.NewCloudMonitoringQueryBuilder()`,
	}
	var buffer strings.Builder
	if input.RefId != "" {

		buffer.WriteString(`RefId(`)
		arg0 := fmt.Sprintf("%#v", input.RefId)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Hide != nil {

		buffer.WriteString(`Hide(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Hide))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.QueryType != nil && *input.QueryType != "" {

		buffer.WriteString(`QueryType(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.QueryType))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.AliasBy != nil && *input.AliasBy != "" {

		buffer.WriteString(`AliasBy(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.AliasBy))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.TimeSeriesList != nil {

		buffer.WriteString(`TimeSeriesList(`)
		arg0 := TimeSeriesListConverter(*input.TimeSeriesList)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.TimeSeriesQuery != nil {

		buffer.WriteString(`TimeSeriesQuery(`)
		arg0 := TimeSeriesQueryConverter(*input.TimeSeriesQuery)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.SloQuery != nil {

		buffer.WriteString(`SloQuery(`)
		arg0 := SLOQueryConverter(*input.SloQuery)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Datasource != nil {

		buffer.WriteString(`Datasource(`)
		arg0 := cog.Dump(*input.Datasource)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.IntervalMs != nil {

		buffer.WriteString(`IntervalMs(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.IntervalMs))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
