// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboard

import (
	"fmt"
	"strings"
)

// TimeOptionConverter accepts a `TimeOption` object and generates the Go code to build this object using builders.
func TimeOptionConverter(input TimeOption) string {
	calls := []string{
		`dashboard.NewTimeOptionBuilder()`,
	}
	var buffer strings.Builder
	if input.Display != "" {

		buffer.WriteString(`Display(`)
		arg0 := fmt.Sprintf("%#v", input.Display)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.From != "" {

		buffer.WriteString(`From(`)
		arg0 := fmt.Sprintf("%#v", input.From)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.To != "" {

		buffer.WriteString(`To(`)
		arg0 := fmt.Sprintf("%#v", input.To)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
