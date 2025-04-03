// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	"fmt"
	"strings"
)

// FilterConverter accepts a `Filter` object and generates the Go code to build this object using builders.
func FilterConverter(input Filter) string {
	calls := []string{
		`elasticsearch.NewFilterBuilder()`,
	}
	var buffer strings.Builder
	if input.Query != "" {

		buffer.WriteString(`Query(`)
		arg0 := fmt.Sprintf("%#v", input.Query)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Label != "" {

		buffer.WriteString(`Label(`)
		arg0 := fmt.Sprintf("%#v", input.Label)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
