// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package rolebinding

import (
	"fmt"
	"strings"
)

func CustomRoleRefConverter(input CustomRoleRef) string {
	calls := []string{
		`rolebinding.NewCustomRoleRefBuilder()`,
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
