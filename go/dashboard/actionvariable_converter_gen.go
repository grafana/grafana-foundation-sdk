// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboard

import (
	"fmt"
	"strings"
)

// ActionVariableConverter accepts a `ActionVariable` object and generates the Go code to build this object using builders.
func ActionVariableConverter(input ActionVariable) string {
	calls := []string{
		`dashboard.NewActionVariableBuilder()`,
	}
	var buffer strings.Builder
	if input.Key != "" {

		buffer.WriteString(`Key(`)
		arg0 := fmt.Sprintf("%#v", input.Key)
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
