// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboard

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// DashboardGraphPanelLegendConverter accepts a `DashboardGraphPanelLegend` object and generates the Go code to build this object using builders.
func DashboardGraphPanelLegendConverter(input DashboardGraphPanelLegend) string {
	calls := []string{
		`dashboard.NewDashboardGraphPanelLegendBuilder()`,
	}
	var buffer strings.Builder
	if input.Show != true {

		buffer.WriteString(`Show(`)
		arg0 := fmt.Sprintf("%#v", input.Show)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Sort != nil && *input.Sort != "" {

		buffer.WriteString(`Sort(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Sort))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.SortDesc != nil {

		buffer.WriteString(`SortDesc(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.SortDesc))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
