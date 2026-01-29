// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"fmt"
	"strings"
)

// TabRepeatOptionsConverter accepts a `TabRepeatOptions` object and generates the Go code to build this object using builders.
func TabRepeatOptionsConverter(input TabRepeatOptions) string {
	calls := []string{
		`dashboardv2beta1.NewTabRepeatOptionsBuilder()`,
	}
	var buffer strings.Builder
	if input.Value != "" {

		buffer.WriteString(`Value(`)
		arg0 := fmt.Sprintf("%#v", input.Value)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
