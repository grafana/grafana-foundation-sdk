// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package statushistory

import (
	"fmt"
	"strings"

	common "github.com/grafana/grafana-foundation-sdk/go/common"
)

// FieldConfigConverter accepts a `FieldConfig` object and generates the Go code to build this object using builders.
func FieldConfigConverter(input FieldConfig) string {
	calls := []string{
		`statushistory.NewFieldConfigBuilder()`,
	}
	var buffer strings.Builder
	if input.LineWidth != nil && *input.LineWidth != 1 {

		buffer.WriteString(`LineWidth(`)
		arg0 := fmt.Sprintf("%#v", *input.LineWidth)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.HideFrom != nil {

		buffer.WriteString(`HideFrom(`)
		arg0 := common.HideSeriesConfigConverter(*input.HideFrom)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.FillOpacity != nil && *input.FillOpacity != 70 {

		buffer.WriteString(`FillOpacity(`)
		arg0 := fmt.Sprintf("%#v", *input.FillOpacity)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
