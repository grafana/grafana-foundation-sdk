// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package cloudwatch

import (
	"fmt"
	"strings"

	dashboardv2beta1 "github.com/grafana/grafana-foundation-sdk/go/dashboardv2beta1"
)

// QueryConverter accepts a `Query` object and generates the Go code to build this object using builders.
func QueryConverter(input dashboardv2beta1.DataQueryKind) string {
	calls := []string{
		`cloudwatch.NewQueryBuilder()`,
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
	if input.Spec != nil && input.Spec.(*CloudWatchQuery).CloudWatchMetricsQuery != nil {

		buffer.WriteString(`CloudWatchMetricsQuery(`)
		arg0 := CloudWatchMetricsQueryConverter(*input.Spec.(*CloudWatchQuery).CloudWatchMetricsQuery)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*CloudWatchQuery).CloudWatchLogsQuery != nil {

		buffer.WriteString(`CloudWatchLogsQuery(`)
		arg0 := CloudWatchLogsQueryConverter(*input.Spec.(*CloudWatchQuery).CloudWatchLogsQuery)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*CloudWatchQuery).CloudWatchAnnotationQuery != nil {

		buffer.WriteString(`CloudWatchAnnotationQuery(`)
		arg0 := CloudWatchAnnotationQueryConverter(*input.Spec.(*CloudWatchQuery).CloudWatchAnnotationQuery)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
