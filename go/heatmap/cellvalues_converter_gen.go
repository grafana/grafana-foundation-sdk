// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package heatmap

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// CellValuesConverter accepts a `CellValues` object and generates the Go code to build this object using builders.
func CellValuesConverter(input CellValues) string {
	calls := []string{
		`heatmap.NewCellValuesBuilder()`,
	}
	var buffer strings.Builder
	if input.Unit != nil && *input.Unit != "" {

		buffer.WriteString(`Unit(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Unit))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Decimals != nil {

		buffer.WriteString(`Decimals(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Decimals))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
