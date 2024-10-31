// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package googlecloudmonitoring

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// LegacyCloudMonitoringAnnotationQueryConverter accepts a `LegacyCloudMonitoringAnnotationQuery` object and generates the Go code to build this object using builders.
func LegacyCloudMonitoringAnnotationQueryConverter(input LegacyCloudMonitoringAnnotationQuery) string {
	calls := []string{
		`googlecloudmonitoring.NewLegacyCloudMonitoringAnnotationQueryBuilder()`,
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
	if input.MetricType != "" {

		buffer.WriteString(`MetricType(`)
		arg0 := fmt.Sprintf("%#v", input.MetricType)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.RefId != "" {

		buffer.WriteString(`RefId(`)
		arg0 := fmt.Sprintf("%#v", input.RefId)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Filters != nil && len(input.Filters) >= 1 {

		buffer.WriteString(`Filters(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Filters {
			tmpfiltersarg1 := fmt.Sprintf("%#v", arg1)
			tmparg0 = append(tmparg0, tmpfiltersarg1)
		}
		arg0 := "[]string{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	{
		buffer.WriteString(`MetricKind(`)
		arg0 := cog.Dump(input.MetricKind)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.ValueType != "" {

		buffer.WriteString(`ValueType(`)
		arg0 := fmt.Sprintf("%#v", input.ValueType)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Title != "" {

		buffer.WriteString(`Title(`)
		arg0 := fmt.Sprintf("%#v", input.Title)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Text != "" {

		buffer.WriteString(`Text(`)
		arg0 := fmt.Sprintf("%#v", input.Text)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
