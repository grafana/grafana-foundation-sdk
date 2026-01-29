// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package tempo

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	dashboardv2beta1 "github.com/grafana/grafana-foundation-sdk/go/dashboardv2beta1"
)

// QueryConverter accepts a `Query` object and generates the Go code to build this object using builders.
func QueryConverter(input dashboardv2beta1.DataQueryKind) string {
	calls := []string{
		`tempo.NewQueryBuilder()`,
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
	if input.Spec != nil && input.Spec.(*TempoQuery).RefId != nil && *input.Spec.(*TempoQuery).RefId != "" {

		buffer.WriteString(`RefId(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*TempoQuery).RefId)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*TempoQuery).Hide != nil {

		buffer.WriteString(`Hide(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*TempoQuery).Hide)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*TempoQuery).QueryType != nil && *input.Spec.(*TempoQuery).QueryType != "" {

		buffer.WriteString(`QueryType(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*TempoQuery).QueryType)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*TempoQuery).Query != nil && *input.Spec.(*TempoQuery).Query != "" {

		buffer.WriteString(`Query(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*TempoQuery).Query)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*TempoQuery).Search != nil && *input.Spec.(*TempoQuery).Search != "" {

		buffer.WriteString(`Search(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*TempoQuery).Search)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*TempoQuery).ServiceName != nil && *input.Spec.(*TempoQuery).ServiceName != "" {

		buffer.WriteString(`ServiceName(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*TempoQuery).ServiceName)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*TempoQuery).SpanName != nil && *input.Spec.(*TempoQuery).SpanName != "" {

		buffer.WriteString(`SpanName(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*TempoQuery).SpanName)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*TempoQuery).MinDuration != nil && *input.Spec.(*TempoQuery).MinDuration != "" {

		buffer.WriteString(`MinDuration(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*TempoQuery).MinDuration)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*TempoQuery).MaxDuration != nil && *input.Spec.(*TempoQuery).MaxDuration != "" {

		buffer.WriteString(`MaxDuration(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*TempoQuery).MaxDuration)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*TempoQuery).ServiceMapQuery != nil {

		buffer.WriteString(`ServiceMapQuery(`)
		arg0 := cog.Dump(*input.Spec.(*TempoQuery).ServiceMapQuery)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*TempoQuery).ServiceMapIncludeNamespace != nil {

		buffer.WriteString(`ServiceMapIncludeNamespace(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*TempoQuery).ServiceMapIncludeNamespace)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*TempoQuery).Limit != nil {

		buffer.WriteString(`Limit(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*TempoQuery).Limit)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*TempoQuery).Spss != nil {

		buffer.WriteString(`Spss(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*TempoQuery).Spss)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*TempoQuery).Filters != nil && len(input.Spec.(*TempoQuery).Filters) >= 1 {

		buffer.WriteString(`Filters(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Spec.(*TempoQuery).Filters {
			tmpfiltersarg1 := TraceqlFilterConverter(arg1)
			tmparg0 = append(tmparg0, tmpfiltersarg1)
		}
		arg0 := "[]cog.Builder[tempo.TraceqlFilter]{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*TempoQuery).GroupBy != nil && len(input.Spec.(*TempoQuery).GroupBy) >= 1 {

		buffer.WriteString(`GroupBy(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Spec.(*TempoQuery).GroupBy {
			tmpgroupByarg1 := TraceqlFilterConverter(arg1)
			tmparg0 = append(tmparg0, tmpgroupByarg1)
		}
		arg0 := "[]cog.Builder[tempo.TraceqlFilter]{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*TempoQuery).TableType != nil {

		buffer.WriteString(`TableType(`)
		arg0 := cog.Dump(*input.Spec.(*TempoQuery).TableType)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*TempoQuery).Step != nil && *input.Spec.(*TempoQuery).Step != "" {

		buffer.WriteString(`Step(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*TempoQuery).Step)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*TempoQuery).Exemplars != nil {

		buffer.WriteString(`Exemplars(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*TempoQuery).Exemplars)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*TempoQuery).Datasource != nil {

		buffer.WriteString(`OldDatasource(`)
		arg0 := cog.Dump(*input.Spec.(*TempoQuery).Datasource)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*TempoQuery).MetricsQueryType != nil {

		buffer.WriteString(`MetricsQueryType(`)
		arg0 := cog.Dump(*input.Spec.(*TempoQuery).MetricsQueryType)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
