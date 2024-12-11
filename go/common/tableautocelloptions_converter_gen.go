// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package common

import (
	"fmt"
	"strings"
)

// TableAutoCellOptionsConverter accepts a `TableAutoCellOptions` object and generates the Go code to build this object using builders.
func TableAutoCellOptionsConverter(input TableAutoCellOptions) string {
	calls := []string{
		`common.NewTableAutoCellOptionsBuilder()`,
	}
	var buffer strings.Builder
	if input.WrapText != nil {

		buffer.WriteString(`WrapText(`)
		arg0 := fmt.Sprintf("%#v", *input.WrapText)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
