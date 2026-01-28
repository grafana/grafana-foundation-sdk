// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package text

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// OptionsConverter accepts a `Options` object and generates the Go code to build this object using builders.
func OptionsConverter(input Options) string {
	calls := []string{
		`text.NewOptionsBuilder()`,
	}
	var buffer strings.Builder

	{
		buffer.WriteString(`Mode(`)
		arg0 := cog.Dump(input.Mode)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.Code != nil {

		buffer.WriteString(`Code(`)
		arg0 := CodeOptionsConverter(*input.Code)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Content != "" && input.Content != "# Title\n\nFor markdown syntax help: [commonmark.org/help](https://commonmark.org/help/)" {

		buffer.WriteString(`Content(`)
		arg0 := fmt.Sprintf("%#v", input.Content)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
