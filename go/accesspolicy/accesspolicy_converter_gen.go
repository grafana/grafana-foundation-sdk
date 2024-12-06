// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package accesspolicy

import (
	"strings"
)

// AccessPolicyConverter accepts a `AccessPolicy` object and generates the Go code to build this object using builders.
func AccessPolicyConverter(input AccessPolicy) string {
	calls := []string{
		`accesspolicy.NewAccessPolicyBuilder()`,
	}
	var buffer strings.Builder

	{
		buffer.WriteString(`Scope(`)
		arg0 := ResourceRefConverter(input.Scope)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	{
		buffer.WriteString(`Role(`)
		arg0 := RoleRefConverter(input.Role)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.Rules != nil && len(input.Rules) >= 1 {
		for _, item := range input.Rules {

			buffer.WriteString(`Rules(`)
			arg0 := AccessRuleConverter(item)
			buffer.WriteString(arg0)

			buffer.WriteString(")")

			calls = append(calls, buffer.String())
			buffer.Reset()

		}
	}

	return strings.Join(calls, ".\t\n")
}
