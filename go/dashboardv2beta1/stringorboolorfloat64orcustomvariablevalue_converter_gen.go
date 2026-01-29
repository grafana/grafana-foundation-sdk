// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// StringOrBoolOrFloat64OrCustomVariableValueConverter accepts a `StringOrBoolOrFloat64OrCustomVariableValue` object and generates the Go code to build this object using builders.
func StringOrBoolOrFloat64OrCustomVariableValueConverter(input StringOrBoolOrFloat64OrCustomVariableValue) string {
	calls := []string{
		`dashboardv2beta1.NewStringOrBoolOrFloat64OrCustomVariableValueBuilder()`,
	}
	var buffer strings.Builder
	if input.String != nil && *input.String != "" {

		buffer.WriteString(`String(`)
		arg0 := fmt.Sprintf("%#v", *input.String)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Bool != nil {

		buffer.WriteString(`Bool(`)
		arg0 := fmt.Sprintf("%#v", *input.Bool)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Float64 != nil {

		buffer.WriteString(`Float64(`)
		arg0 := fmt.Sprintf("%#v", *input.Float64)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.CustomVariableValue != nil {

		buffer.WriteString(`CustomVariableValue(`)
		arg0 := cog.Dump(*input.CustomVariableValue)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
