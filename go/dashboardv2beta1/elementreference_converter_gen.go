// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"fmt"
	"strings"
)

// ElementReferenceConverter accepts a `ElementReference` object and generates the Go code to build this object using builders.
func ElementReferenceConverter(input ElementReference) string {
	calls := []string{
		`dashboardv2beta1.NewElementReferenceBuilder()`,
	}
	var buffer strings.Builder
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
