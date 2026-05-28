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
		arg0 := dashboardv2beta1.Dashboardv2beta1DataQueryKindDatasourceConverter(*input.Datasource)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).RefId != "" {

		buffer.WriteString(`RefId(`)
		arg0 := fmt.Sprintf("%#v", input.Spec.(*Dataquery).RefId)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).Hide != nil {

		buffer.WriteString(`Hide(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*Dataquery).Hide)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).QueryType != nil && *input.Spec.(*Dataquery).QueryType != "" {

		buffer.WriteString(`QueryType(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*Dataquery).QueryType)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).Query != nil && *input.Spec.(*Dataquery).Query != "" {

		buffer.WriteString(`Query(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*Dataquery).Query)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).Search != nil && *input.Spec.(*Dataquery).Search != "" {

		buffer.WriteString(`Search(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*Dataquery).Search)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).ServiceName != nil && *input.Spec.(*Dataquery).ServiceName != "" {

		buffer.WriteString(`ServiceName(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*Dataquery).ServiceName)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).SpanName != nil && *input.Spec.(*Dataquery).SpanName != "" {

		buffer.WriteString(`SpanName(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*Dataquery).SpanName)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).MinDuration != nil && *input.Spec.(*Dataquery).MinDuration != "" {

		buffer.WriteString(`MinDuration(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*Dataquery).MinDuration)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).MaxDuration != nil && *input.Spec.(*Dataquery).MaxDuration != "" {

		buffer.WriteString(`MaxDuration(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*Dataquery).MaxDuration)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).ServiceMapQuery != nil {

		buffer.WriteString(`ServiceMapQuery(`)
		arg0 := cog.Dump(*input.Spec.(*Dataquery).ServiceMapQuery)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).ServiceMapIncludeNamespace != nil {

		buffer.WriteString(`ServiceMapIncludeNamespace(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*Dataquery).ServiceMapIncludeNamespace)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).Limit != nil {

		buffer.WriteString(`Limit(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*Dataquery).Limit)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).Spss != nil {

		buffer.WriteString(`Spss(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*Dataquery).Spss)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).Filters != nil && len(input.Spec.(*Dataquery).Filters) >= 1 {

		buffer.WriteString(`Filters(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Spec.(*Dataquery).Filters {
			tmpfiltersarg1 := TraceqlFilterConverter(arg1)
			tmparg0 = append(tmparg0, tmpfiltersarg1)
		}
		arg0 := "[]cog.Builder[tempo.TraceqlFilter]{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).GroupBy != nil && len(input.Spec.(*Dataquery).GroupBy) >= 1 {

		buffer.WriteString(`GroupBy(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Spec.(*Dataquery).GroupBy {
			tmpgroupByarg1 := TraceqlFilterConverter(arg1)
			tmparg0 = append(tmparg0, tmpgroupByarg1)
		}
		arg0 := "[]cog.Builder[tempo.TraceqlFilter]{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).TableType != nil {

		buffer.WriteString(`TableType(`)
		arg0 := cog.Dump(*input.Spec.(*Dataquery).TableType)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).Step != nil && *input.Spec.(*Dataquery).Step != "" {

		buffer.WriteString(`Step(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*Dataquery).Step)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).Exemplars != nil {

		buffer.WriteString(`Exemplars(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*Dataquery).Exemplars)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).MetricsQueryType != nil {

		buffer.WriteString(`MetricsQueryType(`)
		arg0 := cog.Dump(*input.Spec.(*Dataquery).MetricsQueryType)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
