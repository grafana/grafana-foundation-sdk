// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"fmt"
	"strings"
)

// V2TransformationKindConverter accepts a `V2TransformationKind` object and generates the Go code to build this object using builders.
func V2TransformationKindConverter(input V2TransformationKind) string {
	calls := []string{
		`dashboardv2beta1.NewV2TransformationKindBuilder()`,
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

	{
		buffer.WriteString(`Spec(`)
		arg0 := V2TransformationSpecConverter(input.Spec)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	return strings.Join(calls, ".\t\n")
}
