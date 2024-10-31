// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package accesspolicy

import (
	"fmt"
	"strings"
)

// ResourceRefConverter accepts a `ResourceRef` object and generates the Go code to build this object using builders.
func ResourceRefConverter(input ResourceRef) string {
	calls := []string{
		`accesspolicy.NewResourceRefBuilder()`,
	}
	var buffer strings.Builder
	if input.Kind != "" {

		buffer.WriteString(`Kind(`)
		arg0 := fmt.Sprintf("%#v", input.Kind)
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

	return strings.Join(calls, ".\t\n")
}
