// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	"fmt"
	"strings"
)

// WorkspacesQueryConverter accepts a `WorkspacesQuery` object and generates the Go code to build this object using builders.
func WorkspacesQueryConverter(input WorkspacesQuery) string {
	calls := []string{
		`azuremonitor.NewWorkspacesQueryBuilder()`,
	}
	var buffer strings.Builder
	if input.RawQuery != nil && *input.RawQuery != "" {

		buffer.WriteString(`RawQuery(`)
		arg0 := fmt.Sprintf("%#v", *input.RawQuery)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Subscription != "" {

		buffer.WriteString(`Subscription(`)
		arg0 := fmt.Sprintf("%#v", input.Subscription)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
