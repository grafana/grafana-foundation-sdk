// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package piechart

import (
	"strings"

	common "github.com/grafana/grafana-foundation-sdk/go/common"
)

// FieldConfigConverter accepts a `FieldConfig` object and generates the Go code to build this object using builders.
func FieldConfigConverter(input FieldConfig) string {
	calls := []string{
		`piechart.NewFieldConfigBuilder()`,
	}
	var buffer strings.Builder
	if input.HideFrom != nil {

		buffer.WriteString(`HideFrom(`)
		arg0 := common.HideSeriesConfigConverter(*input.HideFrom)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
