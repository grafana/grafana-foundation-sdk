// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package preferences

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// PreferencesConverter accepts a `Preferences` object and generates the Go code to build this object using builders.
func PreferencesConverter(input Preferences) string {
	calls := []string{
		`preferences.NewPreferencesBuilder()`,
	}
	var buffer strings.Builder
	if input.HomeDashboardUID != nil && *input.HomeDashboardUID != "" {

		buffer.WriteString(`HomeDashboardUID(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.HomeDashboardUID))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Timezone != nil && *input.Timezone != "" {

		buffer.WriteString(`Timezone(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Timezone))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.WeekStart != nil && *input.WeekStart != "" {

		buffer.WriteString(`WeekStart(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.WeekStart))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Theme != nil && *input.Theme != "" {

		buffer.WriteString(`Theme(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Theme))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Language != nil && *input.Language != "" {

		buffer.WriteString(`Language(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Language))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.QueryHistory != nil {

		buffer.WriteString(`QueryHistory(`)
		arg0 := QueryHistoryPreferenceConverter(*input.QueryHistory)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.CookiePreferences != nil {

		buffer.WriteString(`CookiePreferences(`)
		arg0 := CookiePreferencesConverter(*input.CookiePreferences)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Navbar != nil {

		buffer.WriteString(`Navbar(`)
		arg0 := NavbarPreferenceConverter(*input.Navbar)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
