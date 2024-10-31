// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboard

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// ConstantVariableConverter accepts a `ConstantVariable` object and generates the Go code to build this object using builders.
func ConstantVariableConverter(input VariableModel) string {
	calls := []string{
		`dashboard.NewConstantVariableBuilder(` + fmt.Sprintf("%#v", input.Name) + `)`,
	}
	var buffer strings.Builder
	if input.Id != "" && input.Id != "00000000-0000-0000-0000-000000000000" {

		buffer.WriteString(`Id(`)
		arg0 := fmt.Sprintf("%#v", input.Id)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Name != "" {

		buffer.WriteString(`Name(`)
		arg0 := fmt.Sprintf("%#v", input.Name)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Label != nil && *input.Label != "" {

		buffer.WriteString(`Label(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Label))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Description != nil && *input.Description != "" {

		buffer.WriteString(`Description(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Description))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Query != nil {

		buffer.WriteString(`Value(`)
		arg0 := cog.Dump(*input.Query)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.AllFormat != nil && *input.AllFormat != "" {

		buffer.WriteString(`AllFormat(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.AllFormat))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
