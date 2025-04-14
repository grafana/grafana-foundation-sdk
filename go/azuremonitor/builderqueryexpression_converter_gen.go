// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	"fmt"
	"strings"
)

// BuilderQueryExpressionConverter accepts a `BuilderQueryExpression` object and generates the Go code to build this object using builders.
func BuilderQueryExpressionConverter(input BuilderQueryExpression) string {
	calls := []string{
		`azuremonitor.NewBuilderQueryExpressionBuilder()`,
	}
	var buffer strings.Builder
	if input.From != nil {

		buffer.WriteString(`From(`)
		arg0 := BuilderQueryEditorPropertyExpressionConverter(*input.From)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Columns != nil {

		buffer.WriteString(`Columns(`)
		arg0 := BuilderQueryEditorColumnsExpressionConverter(*input.Columns)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Where != nil {

		buffer.WriteString(`Where(`)
		arg0 := BuilderQueryEditorWhereExpressionArrayConverter(*input.Where)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Reduce != nil {

		buffer.WriteString(`Reduce(`)
		arg0 := BuilderQueryEditorReduceExpressionArrayConverter(*input.Reduce)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.GroupBy != nil {

		buffer.WriteString(`GroupBy(`)
		arg0 := BuilderQueryEditorGroupByExpressionArrayConverter(*input.GroupBy)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Limit != nil {

		buffer.WriteString(`Limit(`)
		arg0 := fmt.Sprintf("%#v", *input.Limit)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.OrderBy != nil {

		buffer.WriteString(`OrderBy(`)
		arg0 := BuilderQueryEditorOrderByExpressionArrayConverter(*input.OrderBy)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.FuzzySearch != nil {

		buffer.WriteString(`FuzzySearch(`)
		arg0 := BuilderQueryEditorWhereExpressionArrayConverter(*input.FuzzySearch)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.TimeFilter != nil {

		buffer.WriteString(`TimeFilter(`)
		arg0 := BuilderQueryEditorWhereExpressionArrayConverter(*input.TimeFilter)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
