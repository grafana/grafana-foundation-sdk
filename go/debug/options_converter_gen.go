// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package debug

import (
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// OptionsConverter accepts a `Options` object and generates the Go code to build this object using builders.
func OptionsConverter(input Options) string {
	calls := []string{
		`debug.NewOptionsBuilder()`,
	}
	var buffer strings.Builder

	{
		buffer.WriteString(`Mode(`)
		arg0 := cog.Dump(input.Mode)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.Counters != nil {

		buffer.WriteString(`Counters(`)
		arg0 := UpdateConfigConverter(*input.Counters)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
