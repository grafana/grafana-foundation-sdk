// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"strings"
)

// TabsLayoutSpecConverter accepts a `TabsLayoutSpec` object and generates the Go code to build this object using builders.
func TabsLayoutSpecConverter(input TabsLayoutSpec) string {
	calls := []string{
		`dashboardv2beta1.NewTabsLayoutSpecBuilder()`,
	}
	var buffer strings.Builder
	if input.Tabs != nil && len(input.Tabs) >= 1 {

		buffer.WriteString(`Tabs(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Tabs {
			tmptabsarg1 := TabsLayoutTabConverter(arg1)
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
