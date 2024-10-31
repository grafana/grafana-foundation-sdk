// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package tempo

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// TempoQueryConverter accepts a `TempoQuery` object and generates the Go code to build this object using builders.
func TempoQueryConverter(input TempoQuery) string {
	calls := []string{
		`tempo.NewTempoQueryBuilder()`,
	}
	var buffer strings.Builder
	if input.RefId != "" {

		buffer.WriteString(`RefId(`)
		arg0 := fmt.Sprintf("%#v", input.RefId)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Hide != nil {

		buffer.WriteString(`Hide(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Hide))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.QueryType != nil && *input.QueryType != "" {

		buffer.WriteString(`QueryType(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.QueryType))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Query != nil && *input.Query != "" {

		buffer.WriteString(`Query(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Query))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Search != nil && *input.Search != "" {

		buffer.WriteString(`Search(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Search))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.ServiceName != nil && *input.ServiceName != "" {

		buffer.WriteString(`ServiceName(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.ServiceName))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.SpanName != nil && *input.SpanName != "" {

		buffer.WriteString(`SpanName(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.SpanName))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.MinDuration != nil && *input.MinDuration != "" {

		buffer.WriteString(`MinDuration(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.MinDuration))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.MaxDuration != nil && *input.MaxDuration != "" {

		buffer.WriteString(`MaxDuration(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.MaxDuration))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.ServiceMapQuery != nil {

		buffer.WriteString(`ServiceMapQuery(`)
		arg0 := cog.Dump(*input.ServiceMapQuery)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.ServiceMapIncludeNamespace != nil {

		buffer.WriteString(`ServiceMapIncludeNamespace(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.ServiceMapIncludeNamespace))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Limit != nil {

		buffer.WriteString(`Limit(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Limit))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spss != nil {

		buffer.WriteString(`Spss(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Spss))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Filters != nil && len(input.Filters) >= 1 {

		buffer.WriteString(`Filters(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Filters {
			tmpfiltersarg1 := TraceqlFilterConverter(arg1)
			tmparg0 = append(tmparg0, tmpfiltersarg1)
		}
		arg0 := "[]cog.Builder[tempo.TraceqlFilter]{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.GroupBy != nil && len(input.GroupBy) >= 1 {

		buffer.WriteString(`GroupBy(`)
		tmparg0 := []string{}
		for _, arg1 := range input.GroupBy {
			tmpgroupByarg1 := TraceqlFilterConverter(arg1)
			tmparg0 = append(tmparg0, tmpgroupByarg1)
		}
		arg0 := "[]cog.Builder[tempo.TraceqlFilter]{" + strings.Join(tmparg0, ",\n") + "}"
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
	if input.TableType != nil {

		buffer.WriteString(`TableType(`)
		arg0 := cog.Dump(*input.TableType)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
