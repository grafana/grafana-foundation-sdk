// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package googlecloudmonitoring

import (
	"fmt"
	"strings"
)

// PromQLQueryConverter accepts a `PromQLQuery` object and generates the Go code to build this object using builders.
func PromQLQueryConverter(input PromQLQuery) string {
	calls := []string{
		`googlecloudmonitoring.NewPromQLQueryBuilder()`,
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
	if input.Expr != "" {

		buffer.WriteString(`Expr(`)
		arg0 := fmt.Sprintf("%#v", input.Expr)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Step != "" {

		buffer.WriteString(`Step(`)
		arg0 := fmt.Sprintf("%#v", input.Step)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
