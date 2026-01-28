// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboard

import (
	"fmt"
	"strings"
)

// DataSourceRefConverter accepts a `DataSourceRef` object and generates the Go code to build this object using builders.
func DataSourceRefConverter(input DataSourceRef) string {
	calls := []string{
		`dashboard.NewDataSourceRefBuilder()`,
	}
	var buffer strings.Builder
	if input.Type != nil && *input.Type != "" {

		buffer.WriteString(`Type(`)
		arg0 := fmt.Sprintf("%#v", *input.Type)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Uid != nil && *input.Uid != "" {

		buffer.WriteString(`Uid(`)
		arg0 := fmt.Sprintf("%#v", *input.Uid)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
