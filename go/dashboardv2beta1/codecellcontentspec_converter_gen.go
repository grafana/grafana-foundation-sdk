// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"fmt"
	"strings"
)

// CodeCellContentSpecConverter accepts a `CodeCellContentSpec` object and generates the Go code to build this object using builders.
func CodeCellContentSpecConverter(input CodeCellContentSpec) string {
	calls := []string{
		`dashboardv2beta1.NewCodeCellContentSpecBuilder()`,
	}
	var buffer strings.Builder
	if input.Language != "" {

		buffer.WriteString(`Language(`)
		arg0 := fmt.Sprintf("%#v", input.Language)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Code != "" {

		buffer.WriteString(`Code(`)
		arg0 := fmt.Sprintf("%#v", input.Code)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Highlight != nil && len(input.Highlight) >= 1 {

		buffer.WriteString(`Highlight(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Highlight {
			tmphighlightarg1 := fmt.Sprintf("%#v", arg1)
			tmparg0 = append(tmparg0, tmphighlightarg1)
		}
		arg0 := "[]int64{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Annotation != nil && *input.Annotation != "" {

		buffer.WriteString(`Annotation(`)
		arg0 := fmt.Sprintf("%#v", *input.Annotation)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
