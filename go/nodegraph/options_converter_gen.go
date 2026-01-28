// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package nodegraph

import (
	"strings"
)

// OptionsConverter accepts a `Options` object and generates the Go code to build this object using builders.
func OptionsConverter(input Options) string {
	calls := []string{
		`nodegraph.NewOptionsBuilder()`,
	}
	var buffer strings.Builder
	if input.Nodes != nil {

		buffer.WriteString(`Nodes(`)
		arg0 := NodeOptionsConverter(*input.Nodes)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Edges != nil {

		buffer.WriteString(`Edges(`)
		arg0 := EdgeOptionsConverter(*input.Edges)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
