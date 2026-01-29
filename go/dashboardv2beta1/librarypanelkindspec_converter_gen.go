// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"fmt"
	"strings"
)

// LibraryPanelKindSpecConverter accepts a `LibraryPanelKindSpec` object and generates the Go code to build this object using builders.
func LibraryPanelKindSpecConverter(input LibraryPanelKindSpec) string {
	calls := []string{
		`dashboardv2beta1.NewLibraryPanelKindSpecBuilder()`,
	}
	var buffer strings.Builder

	{
		buffer.WriteString(`Id(`)
		arg0 := fmt.Sprintf("%#v", input.Id)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.Title != "" {

		buffer.WriteString(`Title(`)
		arg0 := fmt.Sprintf("%#v", input.Title)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	{
		buffer.WriteString(`LibraryPanel(`)
		arg0 := LibraryPanelRefConverter(input.LibraryPanel)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	return strings.Join(calls, ".\t\n")
}
