// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"strings"
)

// NotebookLayoutKindConverter accepts a `NotebookLayoutKind` object and generates the Go code to build this object using builders.
func NotebookLayoutKindConverter(input NotebookLayoutKind) string {
	calls := []string{
		`dashboardv2beta1.NewNotebookLayoutKindBuilder()`,
	}
	var buffer strings.Builder

	{
		buffer.WriteString(`Spec(`)
		arg0 := NotebookLayoutSpecConverter(input.Spec)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	return strings.Join(calls, ".\t\n")
}
