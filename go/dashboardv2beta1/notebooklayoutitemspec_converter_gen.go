// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// NotebookLayoutItemSpecConverter accepts a `NotebookLayoutItemSpec` object and generates the Go code to build this object using builders.
func NotebookLayoutItemSpecConverter(input NotebookLayoutItemSpec) string {
	calls := []string{
		`dashboardv2beta1.NewNotebookLayoutItemSpecBuilder()`,
	}
	var buffer strings.Builder

	{
		buffer.WriteString(`Element(`)
		arg0 := ElementReferenceConverter(input.Element)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	{
		buffer.WriteString(`Source(`)
		arg0 := cog.Dump(input.Source)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.Collapsed != nil {

		buffer.WriteString(`Collapsed(`)
		arg0 := fmt.Sprintf("%#v", *input.Collapsed)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
