// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2

import (
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// PreferencesConverter accepts a `Preferences` object and generates the Go code to build this object using builders.
func PreferencesConverter(input Preferences) string {
	calls := []string{
		`dashboardv2.NewPreferencesBuilder()`,
	}
	var buffer strings.Builder
	if input.Layout != nil {

		buffer.WriteString(`Layout(`)
		arg0 := cog.Dump(*input.Layout)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
