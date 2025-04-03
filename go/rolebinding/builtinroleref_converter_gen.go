// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package rolebinding

import (
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// BuiltinRoleRefConverter accepts a `BuiltinRoleRef` object and generates the Go code to build this object using builders.
func BuiltinRoleRefConverter(input BuiltinRoleRef) string {
	calls := []string{
		`rolebinding.NewBuiltinRoleRefBuilder()`,
	}
	var buffer strings.Builder

	{
		buffer.WriteString(`Name(`)
		arg0 := cog.Dump(input.Name)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	return strings.Join(calls, ".\t\n")
}
