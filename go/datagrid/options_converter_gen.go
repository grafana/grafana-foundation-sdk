// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package datagrid

import (
	"fmt"
	"strings"
)

// OptionsConverter accepts a `Options` object and generates the Go code to build this object using builders.
func OptionsConverter(input Options) string {
	calls := []string{
		`datagrid.NewOptionsBuilder()`,
	}
	var buffer strings.Builder
	if input.SelectedSeries != 0 {

		buffer.WriteString(`SelectedSeries(`)
		arg0 := fmt.Sprintf("%#v", input.SelectedSeries)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
