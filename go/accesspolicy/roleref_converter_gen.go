// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package accesspolicy

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// RoleRefConverter accepts a `RoleRef` object and generates the Go code to build this object using builders.
func RoleRefConverter(input RoleRef) string {
	calls := []string{
		`accesspolicy.NewRoleRefBuilder()`,
	}
	var buffer strings.Builder

	{
		buffer.WriteString(`Kind(`)
		arg0 := cog.Dump(input.Kind)
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
	if input.Xname != "" {

		buffer.WriteString(`Xname(`)
		arg0 := fmt.Sprintf("%#v", input.Xname)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
