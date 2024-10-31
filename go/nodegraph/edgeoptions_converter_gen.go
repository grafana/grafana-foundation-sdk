// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package nodegraph

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// EdgeOptionsConverter accepts a `EdgeOptions` object and generates the Go code to build this object using builders.
func EdgeOptionsConverter(input EdgeOptions) string {
	calls := []string{
		`nodegraph.NewEdgeOptionsBuilder()`,
	}
	var buffer strings.Builder
	if input.MainStatUnit != nil && *input.MainStatUnit != "" {

		buffer.WriteString(`MainStatUnit(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.MainStatUnit))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.SecondaryStatUnit != nil && *input.SecondaryStatUnit != "" {

		buffer.WriteString(`SecondaryStatUnit(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.SecondaryStatUnit))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
