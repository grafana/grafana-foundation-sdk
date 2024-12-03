// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package bigquery

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// SQLExpressionConverter accepts a `SQLExpression` object and generates the Go code to build this object using builders.
func SQLExpressionConverter(input SQLExpression) string {
	calls := []string{
		`bigquery.NewSQLExpressionBuilder()`,
	}
	var buffer strings.Builder
	if input.Columns != nil && len(input.Columns) >= 1 {

		buffer.WriteString(`Columns(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Columns {
			tmpcolumnsarg1 := QueryEditorFunctionExpressionConverter(arg1)
			tmparg0 = append(tmparg0, tmpcolumnsarg1)
		}
		arg0 := "[]cog.Builder[bigquery.QueryEditorFunctionExpression]{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.From != nil && *input.From != "" {

		buffer.WriteString(`From(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.From))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.WhereString != nil && *input.WhereString != "" {

		buffer.WriteString(`WhereString(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.WhereString))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.GroupBy != nil && len(input.GroupBy) >= 1 {

		buffer.WriteString(`GroupBy(`)
		tmparg0 := []string{}
		for _, arg1 := range input.GroupBy {
			tmpgroupByarg1 := QueryEditorGroupByExpressionConverter(arg1)
			tmparg0 = append(tmparg0, tmpgroupByarg1)
		}
		arg0 := "[]cog.Builder[bigquery.QueryEditorGroupByExpression]{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.OrderBy != nil {

		buffer.WriteString(`OrderBy(`)
		arg0 := QueryEditorPropertyExpressionConverter(*input.OrderBy)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.OrderByDirection != nil {

		buffer.WriteString(`OrderByDirection(`)
		arg0 := cog.Dump(*input.OrderByDirection)
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
	if input.Offset != nil {

		buffer.WriteString(`Offset(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Offset))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
