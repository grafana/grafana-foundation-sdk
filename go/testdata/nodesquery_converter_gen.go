// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package testdata

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// NodesQueryConverter accepts a `NodesQuery` object and generates the Go code to build this object using builders.
func NodesQueryConverter(input NodesQuery) string {
	calls := []string{
		`testdata.NewNodesQueryBuilder()`,
	}
	var buffer strings.Builder
	if input.Type != nil {

		buffer.WriteString(`Type(`)
		arg0 := cog.Dump(*input.Type)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Count != nil {

		buffer.WriteString(`Count(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Count))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
