// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package googlecloudmonitoring

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// MetricQueryConverter accepts a `MetricQuery` object and generates the Go code to build this object using builders.
func MetricQueryConverter(input MetricQuery) string {
	calls := []string{
		`googlecloudmonitoring.NewMetricQueryBuilder()`,
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
	if input.PerSeriesAligner != nil && *input.PerSeriesAligner != "" {

		buffer.WriteString(`PerSeriesAligner(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.PerSeriesAligner))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.AlignmentPeriod != nil && *input.AlignmentPeriod != "" {

		buffer.WriteString(`AlignmentPeriod(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.AlignmentPeriod))
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
	if input.EditorMode != "" {

		buffer.WriteString(`EditorMode(`)
		arg0 := fmt.Sprintf("%#v", input.EditorMode)
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
	if input.CrossSeriesReducer != "" {

		buffer.WriteString(`CrossSeriesReducer(`)
		arg0 := fmt.Sprintf("%#v", input.CrossSeriesReducer)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.GroupBys != nil && len(input.GroupBys) >= 1 {

		buffer.WriteString(`GroupBys(`)
		tmparg0 := []string{}
		for _, arg1 := range input.GroupBys {
			tmpgroupBysarg1 := fmt.Sprintf("%#v", arg1)
			tmparg0 = append(tmparg0, tmpgroupBysarg1)
		}
		arg0 := "[]string{" + strings.Join(tmparg0, ",\n") + "}"
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
	if input.MetricKind != nil {

		buffer.WriteString(`MetricKind(`)
		arg0 := cog.Dump(*input.MetricKind)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.ValueType != nil && *input.ValueType != "" {

		buffer.WriteString(`ValueType(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.ValueType))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.View != nil && *input.View != "" {

		buffer.WriteString(`View(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.View))
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
	if input.Preprocessor != nil {

		buffer.WriteString(`Preprocessor(`)
		arg0 := cog.Dump(*input.Preprocessor)
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
