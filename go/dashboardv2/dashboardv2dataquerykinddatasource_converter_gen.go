// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2

import (
	"fmt"
	"strings"
)

// Dashboardv2DataQueryKindDatasourceConverter accepts a `Dashboardv2DataQueryKindDatasource` object and generates the Go code to build this object using builders.
func Dashboardv2DataQueryKindDatasourceConverter(input Dashboardv2DataQueryKindDatasource) string {
	calls := []string{
		`dashboardv2.NewDashboardv2DataQueryKindDatasourceBuilder()`,
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
