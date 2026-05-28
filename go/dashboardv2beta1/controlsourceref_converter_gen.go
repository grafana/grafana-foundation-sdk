// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"fmt"
	"strings"
)

// ControlSourceRefConverter accepts a `ControlSourceRef` object and generates the Go code to build this object using builders.
func ControlSourceRefConverter(input ControlSourceRef) string {
	calls := []string{
		`dashboardv2beta1.NewControlSourceRefBuilder()`,
	}
	var buffer strings.Builder
	if input.Group != "" {

		buffer.WriteString(`Group(`)
		arg0 := fmt.Sprintf("%#v", input.Group)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
