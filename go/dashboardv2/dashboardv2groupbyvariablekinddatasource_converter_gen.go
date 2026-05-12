// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2

import (
	"fmt"
	"strings"
)

// Dashboardv2GroupByVariableKindDatasourceConverter accepts a `Dashboardv2GroupByVariableKindDatasource` object and generates the Go code to build this object using builders.
func Dashboardv2GroupByVariableKindDatasourceConverter(input Dashboardv2GroupByVariableKindDatasource) string {
	calls := []string{
		`dashboardv2.NewDashboardv2GroupByVariableKindDatasourceBuilder()`,
	}
	var buffer strings.Builder
	if input.Name != nil && *input.Name != "" {

		buffer.WriteString(`Name(`)
		arg0 := fmt.Sprintf("%#v", *input.Name)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
