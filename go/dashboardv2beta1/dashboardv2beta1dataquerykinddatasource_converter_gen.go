// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"fmt"
	"strings"
)

// Dashboardv2beta1DataQueryKindDatasourceConverter accepts a `Dashboardv2beta1DataQueryKindDatasource` object and generates the Go code to build this object using builders.
func Dashboardv2beta1DataQueryKindDatasourceConverter(input Dashboardv2beta1DataQueryKindDatasource) string {
	calls := []string{
		`dashboardv2beta1.NewDashboardv2beta1DataQueryKindDatasourceBuilder()`,
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
