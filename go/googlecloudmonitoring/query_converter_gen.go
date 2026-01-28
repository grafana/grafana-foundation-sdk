// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package googlecloudmonitoring

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	dashboardv2beta1 "github.com/grafana/grafana-foundation-sdk/go/dashboardv2beta1"
)

// QueryConverter accepts a `Query` object and generates the Go code to build this object using builders.
func QueryConverter(input dashboardv2beta1.DataQueryKind) string {
	calls := []string{
		`googlecloudmonitoring.NewQueryBuilder()`,
	}
	var buffer strings.Builder
	if input.Version != "" && input.Version != "v0" {

		buffer.WriteString(`Version(`)
		arg0 := fmt.Sprintf("%#v", input.Version)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Datasource != nil {

		buffer.WriteString(`Datasource(`)
		arg0 := dashboardv2beta1.Dashboardv2beta1DataQueryKindDatasourceConverter(*input.Datasource)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*CloudMonitoringQuery).RefId != nil && *input.Spec.(*CloudMonitoringQuery).RefId != "" {

		buffer.WriteString(`RefId(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*CloudMonitoringQuery).RefId)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*CloudMonitoringQuery).Hide != nil {

		buffer.WriteString(`Hide(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*CloudMonitoringQuery).Hide)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*CloudMonitoringQuery).QueryType != nil && *input.Spec.(*CloudMonitoringQuery).QueryType != "" {

		buffer.WriteString(`QueryType(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*CloudMonitoringQuery).QueryType)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*CloudMonitoringQuery).AliasBy != nil && *input.Spec.(*CloudMonitoringQuery).AliasBy != "" {

		buffer.WriteString(`AliasBy(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*CloudMonitoringQuery).AliasBy)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*CloudMonitoringQuery).TimeSeriesList != nil {

		buffer.WriteString(`TimeSeriesList(`)
		arg0 := TimeSeriesListConverter(*input.Spec.(*CloudMonitoringQuery).TimeSeriesList)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*CloudMonitoringQuery).TimeSeriesQuery != nil {

		buffer.WriteString(`TimeSeriesQuery(`)
		arg0 := TimeSeriesQueryConverter(*input.Spec.(*CloudMonitoringQuery).TimeSeriesQuery)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*CloudMonitoringQuery).SloQuery != nil {

		buffer.WriteString(`SloQuery(`)
		arg0 := SLOQueryConverter(*input.Spec.(*CloudMonitoringQuery).SloQuery)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*CloudMonitoringQuery).PromQLQuery != nil {

		buffer.WriteString(`PromQLQuery(`)
		arg0 := PromQLQueryConverter(*input.Spec.(*CloudMonitoringQuery).PromQLQuery)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*CloudMonitoringQuery).Datasource != nil {

		buffer.WriteString(`OldDatasource(`)
		arg0 := cog.Dump(*input.Spec.(*CloudMonitoringQuery).Datasource)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*CloudMonitoringQuery).IntervalMs != nil {

		buffer.WriteString(`IntervalMs(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*CloudMonitoringQuery).IntervalMs)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
