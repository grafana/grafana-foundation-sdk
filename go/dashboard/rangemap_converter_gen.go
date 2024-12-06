// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboard

import (
	"strings"
)

// RangeMapConverter accepts a `RangeMap` object and generates the Go code to build this object using builders.
func RangeMapConverter(input RangeMap) string {
	calls := []string{
		`dashboard.NewRangeMapBuilder()`,
	}
	var buffer strings.Builder

	{
		buffer.WriteString(`Options(`)
		arg0 := DashboardRangeMapOptionsConverter(input.Options)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	return strings.Join(calls, ".\t\n")
}
