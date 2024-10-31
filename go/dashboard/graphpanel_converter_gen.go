// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboard

import (
	"strings"
)

// GraphPanelConverter accepts a `GraphPanel` object and generates the Go code to build this object using builders.
func GraphPanelConverter(input GraphPanel) string {
	calls := []string{
		`dashboard.NewGraphPanelBuilder()`,
	}
	var buffer strings.Builder
	if input.Legend != nil {

		buffer.WriteString(`Legend(`)
		arg0 := DashboardGraphPanelLegendConverter(*input.Legend)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
