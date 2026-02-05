// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package bigquery

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// DataqueryConverter accepts a `Dataquery` object and generates the Go code to build this object using builders.
func DataqueryConverter(input Dataquery) string {
	calls := []string{
		`bigquery.NewDataqueryBuilder()`,
	}
	var buffer strings.Builder
	if input.Dataset != nil && *input.Dataset != "" {

		buffer.WriteString(`Dataset(`)
		arg0 := fmt.Sprintf("%#v", *input.Dataset)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Table != nil && *input.Table != "" {

		buffer.WriteString(`Table(`)
		arg0 := fmt.Sprintf("%#v", *input.Table)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Project != nil && *input.Project != "" {

		buffer.WriteString(`Project(`)
		arg0 := fmt.Sprintf("%#v", *input.Project)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	{
		buffer.WriteString(`Format(`)
		arg0 := cog.Dump(input.Format)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.RawQuery != nil {

		buffer.WriteString(`RawQuery(`)
		arg0 := fmt.Sprintf("%#v", *input.RawQuery)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.RawSql != "" {

		buffer.WriteString(`RawSql(`)
		arg0 := fmt.Sprintf("%#v", input.RawSql)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Location != nil && *input.Location != "" {

		buffer.WriteString(`Location(`)
		arg0 := fmt.Sprintf("%#v", *input.Location)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Partitioned != nil {

		buffer.WriteString(`Partitioned(`)
		arg0 := fmt.Sprintf("%#v", *input.Partitioned)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.PartitionedField != nil && *input.PartitionedField != "" {

		buffer.WriteString(`PartitionedField(`)
		arg0 := fmt.Sprintf("%#v", *input.PartitionedField)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.ConvertToUTC != nil {

		buffer.WriteString(`ConvertToUTC(`)
		arg0 := fmt.Sprintf("%#v", *input.ConvertToUTC)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Sharded != nil {

		buffer.WriteString(`Sharded(`)
		arg0 := fmt.Sprintf("%#v", *input.Sharded)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.QueryPriority != nil {

		buffer.WriteString(`QueryPriority(`)
		arg0 := cog.Dump(*input.QueryPriority)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.TimeShift != nil && *input.TimeShift != "" {

		buffer.WriteString(`TimeShift(`)
		arg0 := fmt.Sprintf("%#v", *input.TimeShift)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.EditorMode != nil {

		buffer.WriteString(`EditorMode(`)
		arg0 := cog.Dump(*input.EditorMode)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Sql != nil {

		buffer.WriteString(`Sql(`)
		arg0 := SQLExpressionConverter(*input.Sql)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.RefId != nil && *input.RefId != "" {

		buffer.WriteString(`RefId(`)
		arg0 := fmt.Sprintf("%#v", *input.RefId)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Hide != nil {

		buffer.WriteString(`Hide(`)
		arg0 := fmt.Sprintf("%#v", *input.Hide)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.QueryType != nil && *input.QueryType != "" {

		buffer.WriteString(`QueryType(`)
		arg0 := fmt.Sprintf("%#v", *input.QueryType)
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

	return strings.Join(calls, ".\t\n")
}
