// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// Dashboardv2FieldConfigSourceOverridesConverter accepts a `Dashboardv2FieldConfigSourceOverrides` object and generates the Go code to build this object using builders.
func Dashboardv2FieldConfigSourceOverridesConverter(input Dashboardv2FieldConfigSourceOverrides) string {
	calls := []string{
		`dashboardv2.NewDashboardv2FieldConfigSourceOverridesBuilder()`,
	}
	var buffer strings.Builder
	if input.SystemRef != nil && *input.SystemRef != "" {

		buffer.WriteString(`SystemRef(`)
		arg0 := fmt.Sprintf("%#v", *input.SystemRef)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

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
		arg0 := "[]dashboardv2.DynamicConfigValue{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
