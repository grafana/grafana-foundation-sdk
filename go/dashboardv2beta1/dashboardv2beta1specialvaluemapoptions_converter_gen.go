// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// Dashboardv2beta1SpecialValueMapOptionsConverter accepts a `Dashboardv2beta1SpecialValueMapOptions` object and generates the Go code to build this object using builders.
func Dashboardv2beta1SpecialValueMapOptionsConverter(input Dashboardv2beta1SpecialValueMapOptions) string {
	calls := []string{
		`dashboardv2beta1.NewDashboardv2beta1SpecialValueMapOptionsBuilder()`,
	}
	var buffer strings.Builder

	{
		buffer.WriteString(`Match(`)
		arg0 := cog.Dump(input.Match)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	{
		buffer.WriteString(`Result(`)
		arg0 := ValueMappingResultConverter(input.Result)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	return strings.Join(calls, ".\t\n")
}
