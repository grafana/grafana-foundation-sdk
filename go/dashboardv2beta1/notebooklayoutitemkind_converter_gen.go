// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"strings"
)

// NotebookLayoutItemKindConverter accepts a `NotebookLayoutItemKind` object and generates the Go code to build this object using builders.
func NotebookLayoutItemKindConverter(input NotebookLayoutItemKind) string {
	calls := []string{
		`dashboardv2beta1.NewNotebookLayoutItemKindBuilder()`,
	}
	var buffer strings.Builder

	{
		buffer.WriteString(`Spec(`)
		arg0 := NotebookLayoutItemSpecConverter(input.Spec)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	return strings.Join(calls, ".\t\n")
}
