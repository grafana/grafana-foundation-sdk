// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package cloudwatch

import (
	"fmt"
	"strings"

	dashboardv2 "github.com/grafana/grafana-foundation-sdk/go/dashboardv2"
)

// QueryV2Converter accepts a `QueryV2` object and generates the Go code to build this object using builders.
func QueryV2Converter(input dashboardv2.DataQueryKind) string {
	calls := []string{
		`cloudwatch.NewQueryV2Builder()`,
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
	if input.Labels != nil {

		buffer.WriteString(`Labels(`)
		arg0 := "map[string]string{"
		for key, arg1 := range input.Labels {
			tmplabelsarg1 := fmt.Sprintf("%#v", arg1)
			arg0 += "\t" + fmt.Sprintf("%#v", key) + ": " + tmplabelsarg1 + ","
		}
		arg0 += "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Datasource != nil {

		buffer.WriteString(`Datasource(`)
		arg0 := dashboardv2.Dashboardv2DataQueryKindDatasourceConverter(*input.Datasource)
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
