// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// DataQueryKindConverter accepts a `DataQueryKind` object and generates the Go code to build this object using builders.
func DataQueryKindConverter(input DataQueryKind) string {
	calls := []string{
		`dashboardv2beta1.NewDataQueryKindBuilder()`,
	}
	var buffer strings.Builder
	if input.Group != "" {

		buffer.WriteString(`Group(`)
		arg0 := fmt.Sprintf("%#v", input.Group)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Version != "" && input.Version != "v0" {

		buffer.WriteString(`Version(`)
		arg0 := fmt.Sprintf("%#v", input.Version)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Datasource != nil {

		buffer.WriteString(`Datasource(`)
		arg0 := Dashboardv2beta1DataQueryKindDatasourceConverter(*input.Datasource)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil {

		buffer.WriteString(`Spec(`)
		arg0 := cog.Dump(input.Spec)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
