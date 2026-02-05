// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package bigquery

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	dashboardv2beta1 "github.com/grafana/grafana-foundation-sdk/go/dashboardv2beta1"
)

// QueryConverter accepts a `Query` object and generates the Go code to build this object using builders.
func QueryConverter(input dashboardv2beta1.DataQueryKind) string {
	calls := []string{
		`bigquery.NewQueryBuilder()`,
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
	if input.Spec != nil && input.Spec.(*Dataquery).Dataset != nil && *input.Spec.(*Dataquery).Dataset != "" {

		buffer.WriteString(`Dataset(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*Dataquery).Dataset)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).Table != nil && *input.Spec.(*Dataquery).Table != "" {

		buffer.WriteString(`Table(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*Dataquery).Table)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).Project != nil && *input.Spec.(*Dataquery).Project != "" {

		buffer.WriteString(`Project(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*Dataquery).Project)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil {

		buffer.WriteString(`Format(`)
		arg0 := cog.Dump(input.Spec.(*Dataquery).Format)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).RawQuery != nil {

		buffer.WriteString(`RawQuery(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*Dataquery).RawQuery)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).RawSql != "" {

		buffer.WriteString(`RawSql(`)
		arg0 := fmt.Sprintf("%#v", input.Spec.(*Dataquery).RawSql)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).Location != nil && *input.Spec.(*Dataquery).Location != "" {

		buffer.WriteString(`Location(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*Dataquery).Location)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).Partitioned != nil {

		buffer.WriteString(`Partitioned(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*Dataquery).Partitioned)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).PartitionedField != nil && *input.Spec.(*Dataquery).PartitionedField != "" {

		buffer.WriteString(`PartitionedField(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*Dataquery).PartitionedField)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).ConvertToUTC != nil {

		buffer.WriteString(`ConvertToUTC(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*Dataquery).ConvertToUTC)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).Sharded != nil {

		buffer.WriteString(`Sharded(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*Dataquery).Sharded)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).QueryPriority != nil {

		buffer.WriteString(`QueryPriority(`)
		arg0 := cog.Dump(*input.Spec.(*Dataquery).QueryPriority)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).TimeShift != nil && *input.Spec.(*Dataquery).TimeShift != "" {

		buffer.WriteString(`TimeShift(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*Dataquery).TimeShift)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).EditorMode != nil {

		buffer.WriteString(`EditorMode(`)
		arg0 := cog.Dump(*input.Spec.(*Dataquery).EditorMode)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).Sql != nil {

		buffer.WriteString(`Sql(`)
		arg0 := SQLExpressionConverter(*input.Spec.(*Dataquery).Sql)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).RefId != nil && *input.Spec.(*Dataquery).RefId != "" {

		buffer.WriteString(`RefId(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*Dataquery).RefId)
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

	return strings.Join(calls, ".\t\n")
}
