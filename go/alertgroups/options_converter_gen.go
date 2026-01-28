// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alertgroups

import (
	"fmt"
	"strings"
)

// OptionsConverter accepts a `Options` object and generates the Go code to build this object using builders.
func OptionsConverter(input Options) string {
	calls := []string{
		`alertgroups.NewOptionsBuilder()`,
	}
	var buffer strings.Builder
	if input.Labels != "" {

		buffer.WriteString(`Labels(`)
		arg0 := fmt.Sprintf("%#v", input.Labels)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Alertmanager != "" {

		buffer.WriteString(`Alertmanager(`)
		arg0 := fmt.Sprintf("%#v", input.Alertmanager)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	{
		buffer.WriteString(`ExpandAll(`)
		arg0 := fmt.Sprintf("%#v", input.ExpandAll)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	return strings.Join(calls, ".\t\n")
}
