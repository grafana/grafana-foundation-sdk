// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package cloudwatch

import (
	"strings"
)

// MetricsQueryOrLogsQueryOrAnnotationQueryConverter accepts a `MetricsQueryOrLogsQueryOrAnnotationQuery` object and generates the Go code to build this object using builders.
func MetricsQueryOrLogsQueryOrAnnotationQueryConverter(input MetricsQueryOrLogsQueryOrAnnotationQuery) string {
	calls := []string{
		`cloudwatch.NewMetricsQueryOrLogsQueryOrAnnotationQueryBuilder()`,
	}
	var buffer strings.Builder
	if input.MetricsQuery != nil {

		buffer.WriteString(`MetricsQuery(`)
		arg0 := MetricsQueryConverter(*input.MetricsQuery)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.LogsQuery != nil {

		buffer.WriteString(`LogsQuery(`)
		arg0 := LogsQueryConverter(*input.LogsQuery)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.AnnotationQuery != nil {

		buffer.WriteString(`AnnotationQuery(`)
		arg0 := AnnotationQueryConverter(*input.AnnotationQuery)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
