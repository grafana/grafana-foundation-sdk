// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package common

import (
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// GraphThresholdsStyleConfigConverter accepts a `GraphThresholdsStyleConfig` object and generates the Go code to build this object using builders.
func GraphThresholdsStyleConfigConverter(input GraphThresholdsStyleConfig) string {
	calls := []string{
		`common.NewGraphThresholdsStyleConfigBuilder()`,
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

	return strings.Join(calls, ".\t\n")
}
