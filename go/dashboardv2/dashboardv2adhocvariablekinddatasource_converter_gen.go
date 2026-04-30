// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2

import (
	"fmt"
	"strings"
)

// Dashboardv2AdhocVariableKindDatasourceConverter accepts a `Dashboardv2AdhocVariableKindDatasource` object and generates the Go code to build this object using builders.
func Dashboardv2AdhocVariableKindDatasourceConverter(input Dashboardv2AdhocVariableKindDatasource) string {
	calls := []string{
		`dashboardv2.NewDashboardv2AdhocVariableKindDatasourceBuilder()`,
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
