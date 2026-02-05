// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	"fmt"
	"strings"
)

// CountConverter accepts a `Count` object and generates the Go code to build this object using builders.
func CountConverter(input Count) string {
	calls := []string{
		`elasticsearch.NewCountBuilder()`,
	}
	var buffer strings.Builder
	if input.Id != "" {

		buffer.WriteString(`Id(`)
		arg0 := fmt.Sprintf("%#v", input.Id)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Hide != nil {

		buffer.WriteString(`Hide(`)
		arg0 := fmt.Sprintf("%#v", *input.Hide)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
