// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboard

import (
	"strings"
)

// RegexMapConverter accepts a `RegexMap` object and generates the Go code to build this object using builders.
func RegexMapConverter(input RegexMap) string {
	calls := []string{
		`dashboard.NewRegexMapBuilder()`,
	}
	var buffer strings.Builder

	{
		buffer.WriteString(`Options(`)
		arg0 := DashboardRegexMapOptionsConverter(input.Options)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	return strings.Join(calls, ".\t\n")
}
