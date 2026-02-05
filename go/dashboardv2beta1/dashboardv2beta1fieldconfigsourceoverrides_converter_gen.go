// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// Dashboardv2beta1FieldConfigSourceOverridesConverter accepts a `Dashboardv2beta1FieldConfigSourceOverrides` object and generates the Go code to build this object using builders.
func Dashboardv2beta1FieldConfigSourceOverridesConverter(input Dashboardv2beta1FieldConfigSourceOverrides) string {
	calls := []string{
		`dashboardv2beta1.NewDashboardv2beta1FieldConfigSourceOverridesBuilder()`,
	}
	var buffer strings.Builder

	{
		buffer.WriteString(`Matcher(`)
		arg0 := cog.Dump(input.Matcher)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.Properties != nil && len(input.Properties) >= 1 {

		buffer.WriteString(`Properties(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Properties {
			tmppropertiesarg1 := cog.Dump(arg1)
			tmparg0 = append(tmparg0, tmppropertiesarg1)
		}
		arg0 := "[]dashboardv2beta1.DynamicConfigValue{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
