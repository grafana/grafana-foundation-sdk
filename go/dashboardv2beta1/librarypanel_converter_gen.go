// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"fmt"
	"strings"
)

// LibraryPanelConverter accepts a `LibraryPanel` object and generates the Go code to build this object using builders.
func LibraryPanelConverter(input LibraryPanelKind) string {
	calls := []string{
		`dashboardv2beta1.NewLibraryPanelBuilder()`,
	}
	var buffer strings.Builder

	{
		buffer.WriteString(`Id(`)
		arg0 := fmt.Sprintf("%#v", input.Spec.Id)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.Spec.Title != "" {

		buffer.WriteString(`Title(`)
		arg0 := fmt.Sprintf("%#v", input.Spec.Title)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	{
		buffer.WriteString(`LibraryPanel(`)
		arg0 := LibraryPanelRefConverter(input.Spec.LibraryPanel)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	return strings.Join(calls, ".\t\n")
}
