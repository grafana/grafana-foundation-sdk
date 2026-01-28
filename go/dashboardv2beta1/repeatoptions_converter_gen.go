// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// RepeatOptionsConverter accepts a `RepeatOptions` object and generates the Go code to build this object using builders.
func RepeatOptionsConverter(input RepeatOptions) string {
	calls := []string{
		`dashboardv2beta1.NewRepeatOptionsBuilder()`,
	}
	var buffer strings.Builder
	if input.Value != "" {

		buffer.WriteString(`Value(`)
		arg0 := fmt.Sprintf("%#v", input.Value)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Direction != nil {

		buffer.WriteString(`Direction(`)
		arg0 := cog.Dump(*input.Direction)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.MaxPerRow != nil {

		buffer.WriteString(`MaxPerRow(`)
		arg0 := fmt.Sprintf("%#v", *input.MaxPerRow)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
