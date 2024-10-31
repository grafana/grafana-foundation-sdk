// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package common

import (
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// TableColoredBackgroundCellOptionsConverter accepts a `TableColoredBackgroundCellOptions` object and generates the Go code to build this object using builders.
func TableColoredBackgroundCellOptionsConverter(input TableColoredBackgroundCellOptions) string {
	calls := []string{
		`common.NewTableColoredBackgroundCellOptionsBuilder()`,
	}
	var buffer strings.Builder
	if input.Mode != nil {

		buffer.WriteString(`Mode(`)
		arg0 := cog.Dump(*input.Mode)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
