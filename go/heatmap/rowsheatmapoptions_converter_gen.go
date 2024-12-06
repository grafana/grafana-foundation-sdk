// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package heatmap

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// RowsHeatmapOptionsConverter accepts a `RowsHeatmapOptions` object and generates the Go code to build this object using builders.
func RowsHeatmapOptionsConverter(input RowsHeatmapOptions) string {
	calls := []string{
		`heatmap.NewRowsHeatmapOptionsBuilder()`,
	}
	var buffer strings.Builder
	if input.Value != nil && *input.Value != "" {

		buffer.WriteString(`Value(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Value))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Layout != nil {

		buffer.WriteString(`Layout(`)
		arg0 := cog.Dump(*input.Layout)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
