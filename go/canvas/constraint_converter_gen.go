// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package canvas

import (
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// ConstraintConverter accepts a `Constraint` object and generates the Go code to build this object using builders.
func ConstraintConverter(input Constraint) string {
	calls := []string{
		`canvas.NewConstraintBuilder()`,
	}
	var buffer strings.Builder
	if input.Horizontal != nil {

		buffer.WriteString(`Horizontal(`)
		arg0 := cog.Dump(*input.Horizontal)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Vertical != nil {

		buffer.WriteString(`Vertical(`)
		arg0 := cog.Dump(*input.Vertical)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
