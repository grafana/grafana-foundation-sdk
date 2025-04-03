// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package rolebinding

import (
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// RoleBindingConverter accepts a `RoleBinding` object and generates the Go code to build this object using builders.
func RoleBindingConverter(input RoleBinding) string {
	calls := []string{
		`rolebinding.NewRoleBindingBuilder()`,
	}
	var buffer strings.Builder

	{
		buffer.WriteString(`Role(`)
		arg0 := cog.Dump(input.Role)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	{
		buffer.WriteString(`Subject(`)
		arg0 := RoleBindingSubjectConverter(input.Subject)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	return strings.Join(calls, ".\t\n")
}
