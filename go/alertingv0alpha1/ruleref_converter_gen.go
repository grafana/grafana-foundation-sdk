// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alertingv0alpha1

import (
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// RuleRefConverter accepts a `RuleRef` object and generates the Go code to build this object using builders.
func RuleRefConverter(input RuleRef) string {
	calls := []string{
		`alertingv0alpha1.NewRuleRefBuilder()`,
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
