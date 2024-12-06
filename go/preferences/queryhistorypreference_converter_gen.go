// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package preferences

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// QueryHistoryPreferenceConverter accepts a `QueryHistoryPreference` object and generates the Go code to build this object using builders.
func QueryHistoryPreferenceConverter(input QueryHistoryPreference) string {
	calls := []string{
		`preferences.NewQueryHistoryPreferenceBuilder()`,
	}
	var buffer strings.Builder
	if input.HomeTab != nil && *input.HomeTab != "" {

		buffer.WriteString(`HomeTab(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.HomeTab))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
