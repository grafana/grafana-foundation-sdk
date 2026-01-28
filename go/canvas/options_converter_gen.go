// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package canvas

import (
	"fmt"
	"strings"
)

// OptionsConverter accepts a `Options` object and generates the Go code to build this object using builders.
func OptionsConverter(input Options) string {
	calls := []string{
		`canvas.NewOptionsBuilder()`,
	}
	var buffer strings.Builder
	if input.InlineEditing != true {

		buffer.WriteString(`InlineEditing(`)
		arg0 := fmt.Sprintf("%#v", input.InlineEditing)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.ShowAdvancedTypes != true {

		buffer.WriteString(`ShowAdvancedTypes(`)
		arg0 := fmt.Sprintf("%#v", input.ShowAdvancedTypes)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	{
		buffer.WriteString(`Root(`)
		arg0 := CanvasOptionsRootConverter(input.Root)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	return strings.Join(calls, ".\t\n")
}
