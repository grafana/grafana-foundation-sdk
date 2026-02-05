// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package preferences

import (
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// CookiePreferencesConverter accepts a `CookiePreferences` object and generates the Go code to build this object using builders.
func CookiePreferencesConverter(input CookiePreferences) string {
	calls := []string{
		`preferences.NewCookiePreferencesBuilder()`,
	}
	var buffer strings.Builder
	if input.Analytics != nil {

		buffer.WriteString(`Analytics(`)
		arg0 := cog.Dump(input.Analytics)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Performance != nil {

		buffer.WriteString(`Performance(`)
		arg0 := cog.Dump(input.Performance)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Functional != nil {

		buffer.WriteString(`Functional(`)
		arg0 := cog.Dump(input.Functional)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
