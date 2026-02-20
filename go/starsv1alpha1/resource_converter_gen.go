// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package starsv1alpha1

import (
	"fmt"
	"strings"
)

// ResourceConverter accepts a `Resource` object and generates the Go code to build this object using builders.
func ResourceConverter(input Resource) string {
	calls := []string{
		`starsv1alpha1.NewResourceBuilder()`,
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
	if input.Kind != "" {

		buffer.WriteString(`Kind(`)
		arg0 := fmt.Sprintf("%#v", input.Kind)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Names != nil && len(input.Names) >= 1 {

		buffer.WriteString(`Names(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Names {
			tmpnamesarg1 := fmt.Sprintf("%#v", arg1)
			tmparg0 = append(tmparg0, tmpnamesarg1)
		}
		arg0 := "[]string{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
