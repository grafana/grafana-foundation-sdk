// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package nodegraph

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// ArcOptionConverter accepts a `ArcOption` object and generates the Go code to build this object using builders.
func ArcOptionConverter(input ArcOption) string {
	calls := []string{
		`nodegraph.NewArcOptionBuilder()`,
	}
	var buffer strings.Builder
	if input.Field != nil && *input.Field != "" {

		buffer.WriteString(`Field(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Field))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Color != nil && *input.Color != "" {

		buffer.WriteString(`Color(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Color))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
