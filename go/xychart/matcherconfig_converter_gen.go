// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package xychart

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// MatcherConfigConverter accepts a `MatcherConfig` object and generates the Go code to build this object using builders.
func MatcherConfigConverter(input MatcherConfig) string {
	calls := []string{
		`xychart.NewMatcherConfigBuilder()`,
	}
	var buffer strings.Builder
	if input.Id != "" {

		buffer.WriteString(`Id(`)
		arg0 := fmt.Sprintf("%#v", input.Id)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Options != nil {

		buffer.WriteString(`Options(`)
		arg0 := cog.Dump(input.Options)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
