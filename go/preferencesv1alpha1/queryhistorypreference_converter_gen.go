// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package preferencesv1alpha1

import (
	"fmt"
	"strings"
)

// QueryHistoryPreferenceConverter accepts a `QueryHistoryPreference` object and generates the Go code to build this object using builders.
func QueryHistoryPreferenceConverter(input QueryHistoryPreference) string {
	calls := []string{
		`preferencesv1alpha1.NewQueryHistoryPreferenceBuilder()`,
	}
	var buffer strings.Builder
	if input.HomeTab != nil && *input.HomeTab != "" {

		buffer.WriteString(`HomeTab(`)
		arg0 := fmt.Sprintf("%#v", *input.HomeTab)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
