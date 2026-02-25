// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"strings"
)

// TabsConverter accepts a `Tabs` object and generates the Go code to build this object using builders.
func TabsConverter(input TabsLayoutKind) string {
	calls := []string{
		`dashboardv2beta1.NewTabsBuilder()`,
	}
	var buffer strings.Builder
	if input.Spec.Tabs != nil && len(input.Spec.Tabs) >= 1 {

		buffer.WriteString(`Tabs(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Spec.Tabs {
			tmptabsarg1 := TabConverter(arg1)
			tmparg0 = append(tmparg0, tmptabsarg1)
		}
		arg0 := "[]cog.Builder[dashboardv2beta1.TabsLayoutTabKind]{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
