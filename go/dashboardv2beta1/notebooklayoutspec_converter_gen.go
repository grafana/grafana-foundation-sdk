// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"strings"
)

// NotebookLayoutSpecConverter accepts a `NotebookLayoutSpec` object and generates the Go code to build this object using builders.
func NotebookLayoutSpecConverter(input NotebookLayoutSpec) string {
	calls := []string{
		`dashboardv2beta1.NewNotebookLayoutSpecBuilder()`,
	}
	var buffer strings.Builder
	if input.Cells != nil && len(input.Cells) >= 1 {

		buffer.WriteString(`Cells(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Cells {
			tmpcellsarg1 := NotebookLayoutItemKindConverter(arg1)
			tmparg0 = append(tmparg0, tmpcellsarg1)
		}
		arg0 := "[]cog.Builder[dashboardv2beta1.NotebookLayoutItemKind]{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
