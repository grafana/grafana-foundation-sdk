// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"fmt"
	"strings"
)

// Dashboardv2beta1AdhocVariableKindDatasourceConverter accepts a `Dashboardv2beta1AdhocVariableKindDatasource` object and generates the Go code to build this object using builders.
func Dashboardv2beta1AdhocVariableKindDatasourceConverter(input Dashboardv2beta1AdhocVariableKindDatasource) string {
	calls := []string{
		`dashboardv2beta1.NewDashboardv2beta1AdhocVariableKindDatasourceBuilder()`,
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
