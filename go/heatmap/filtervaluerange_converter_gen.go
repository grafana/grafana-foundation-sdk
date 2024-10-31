// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package heatmap

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// FilterValueRangeConverter accepts a `FilterValueRange` object and generates the Go code to build this object using builders.
func FilterValueRangeConverter(input FilterValueRange) string {
	calls := []string{
		`heatmap.NewFilterValueRangeBuilder()`,
	}
	var buffer strings.Builder
	if input.Le != nil {

		buffer.WriteString(`Le(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Le))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Ge != nil {

		buffer.WriteString(`Ge(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Ge))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
