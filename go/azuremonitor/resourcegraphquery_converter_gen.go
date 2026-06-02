// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	"fmt"
	"strings"
)

// ResourceGraphQueryConverter accepts a `ResourceGraphQuery` object and generates the Go code to build this object using builders.
func ResourceGraphQueryConverter(input ResourceGraphQuery) string {
	calls := []string{
		`azuremonitor.NewResourceGraphQueryBuilder()`,
	}
	var buffer strings.Builder
	if input.Query != nil && *input.Query != "" {

		buffer.WriteString(`Query(`)
		arg0 := fmt.Sprintf("%#v", *input.Query)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.ResultFormat != nil && *input.ResultFormat != "" {

		buffer.WriteString(`ResultFormat(`)
		arg0 := fmt.Sprintf("%#v", *input.ResultFormat)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
