// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package debug

import (
	"fmt"
	"strings"
)

// UpdateConfigConverter accepts a `UpdateConfig` object and generates the Go code to build this object using builders.
func UpdateConfigConverter(input UpdateConfig) string {
	calls := []string{
		`debug.NewUpdateConfigBuilder()`,
	}
	var buffer strings.Builder

	{
		buffer.WriteString(`Render(`)
		arg0 := fmt.Sprintf("%#v", input.Render)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	{
		buffer.WriteString(`DataChanged(`)
		arg0 := fmt.Sprintf("%#v", input.DataChanged)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	{
		buffer.WriteString(`SchemaChanged(`)
		arg0 := fmt.Sprintf("%#v", input.SchemaChanged)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	return strings.Join(calls, ".\t\n")
}
