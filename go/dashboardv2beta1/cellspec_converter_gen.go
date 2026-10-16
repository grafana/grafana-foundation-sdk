// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// CellSpecConverter accepts a `CellSpec` object and generates the Go code to build this object using builders.
func CellSpecConverter(input CellSpec) string {
	calls := []string{
		`dashboardv2beta1.NewCellSpecBuilder()`,
	}
	var buffer strings.Builder

	{
		buffer.WriteString(`Content(`)
		arg0 := cog.Dump(input.Content)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	return strings.Join(calls, ".\t\n")
}
