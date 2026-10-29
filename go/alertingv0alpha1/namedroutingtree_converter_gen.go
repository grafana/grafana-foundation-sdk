// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alertingv0alpha1

import (
	"fmt"
	"strings"
)

// NamedRoutingTreeConverter accepts a `NamedRoutingTree` object and generates the Go code to build this object using builders.
func NamedRoutingTreeConverter(input NamedRoutingTree) string {
	calls := []string{
		`alertingv0alpha1.NewNamedRoutingTreeBuilder()`,
	}
	var buffer strings.Builder
	if input.RoutingTree != "" {

		buffer.WriteString(`RoutingTree(`)
		arg0 := fmt.Sprintf("%#v", input.RoutingTree)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
