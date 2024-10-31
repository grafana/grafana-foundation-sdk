// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package accesspolicy

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// AccessRuleConverter accepts a `AccessRule` object and generates the Go code to build this object using builders.
func AccessRuleConverter(input AccessRule) string {
	calls := []string{
		`accesspolicy.NewAccessRuleBuilder()`,
	}
	var buffer strings.Builder
	if input.Kind != "" && input.Kind != "*" {

		buffer.WriteString(`Kind(`)
		arg0 := fmt.Sprintf("%#v", input.Kind)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Verb != "" {

		buffer.WriteString(`Verb(`)
		arg0 := fmt.Sprintf("%#v", input.Verb)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Target != nil && *input.Target != "" {

		buffer.WriteString(`Target(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Target))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
