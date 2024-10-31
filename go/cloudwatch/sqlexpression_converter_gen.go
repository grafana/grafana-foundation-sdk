// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package cloudwatch

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// SQLExpressionConverter accepts a `SQLExpression` object and generates the Go code to build this object using builders.
func SQLExpressionConverter(input SQLExpression) string {
	calls := []string{
		`cloudwatch.NewSQLExpressionBuilder()`,
	}
	var buffer strings.Builder
	if input.Select != nil {

		buffer.WriteString(`Select(`)
		arg0 := QueryEditorFunctionExpressionConverter(*input.Select)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.From != nil {

		buffer.WriteString(`From(`)
		arg0 := cog.Dump(*input.From)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Where != nil {

		buffer.WriteString(`Where(`)
		arg0 := QueryEditorArrayExpressionConverter(*input.Where)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.GroupBy != nil {

		buffer.WriteString(`GroupBy(`)
		arg0 := QueryEditorArrayExpressionConverter(*input.GroupBy)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.OrderBy != nil {

		buffer.WriteString(`OrderBy(`)
		arg0 := QueryEditorFunctionExpressionConverter(*input.OrderBy)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.OrderByDirection != nil && *input.OrderByDirection != "" {

		buffer.WriteString(`OrderByDirection(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.OrderByDirection))
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

	return strings.Join(calls, ".\t\n")
}
