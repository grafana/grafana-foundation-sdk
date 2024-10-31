// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package nodegraph

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// NodeOptionsConverter accepts a `NodeOptions` object and generates the Go code to build this object using builders.
func NodeOptionsConverter(input NodeOptions) string {
	calls := []string{
		`nodegraph.NewNodeOptionsBuilder()`,
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
	if input.Arcs != nil && len(input.Arcs) >= 1 {

		buffer.WriteString(`Arcs(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Arcs {
			tmparcsarg1 := ArcOptionConverter(arg1)
			tmparg0 = append(tmparg0, tmparcsarg1)
		}
		arg0 := "[]cog.Builder[nodegraph.ArcOption]{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
