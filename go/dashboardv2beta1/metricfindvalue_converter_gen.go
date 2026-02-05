// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// MetricFindValueConverter accepts a `MetricFindValue` object and generates the Go code to build this object using builders.
func MetricFindValueConverter(input MetricFindValue) string {
	calls := []string{
		`dashboardv2beta1.NewMetricFindValueBuilder()`,
	}
	var buffer strings.Builder
	if input.Text != "" {

		buffer.WriteString(`Text(`)
		arg0 := fmt.Sprintf("%#v", input.Text)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Value != nil {

		buffer.WriteString(`Value(`)
		arg0 := cog.Dump(*input.Value)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Group != nil && *input.Group != "" {

		buffer.WriteString(`Group(`)
		arg0 := fmt.Sprintf("%#v", *input.Group)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Expandable != nil {

		buffer.WriteString(`Expandable(`)
		arg0 := fmt.Sprintf("%#v", *input.Expandable)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
