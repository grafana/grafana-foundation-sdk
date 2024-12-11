// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package common

import (
	"fmt"
	"strings"
)

// TableImageCellOptionsConverter accepts a `TableImageCellOptions` object and generates the Go code to build this object using builders.
func TableImageCellOptionsConverter(input TableImageCellOptions) string {
	calls := []string{
		`common.NewTableImageCellOptionsBuilder()`,
	}
	var buffer strings.Builder
	if input.Alt != nil && *input.Alt != "" {

		buffer.WriteString(`Alt(`)
		arg0 := fmt.Sprintf("%#v", *input.Alt)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Title != nil && *input.Title != "" {

		buffer.WriteString(`Title(`)
		arg0 := fmt.Sprintf("%#v", *input.Title)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
