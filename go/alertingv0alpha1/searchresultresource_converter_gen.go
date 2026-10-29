// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alertingv0alpha1

import (
	"fmt"
	"strings"
)

// SearchResultResourceConverter accepts a `SearchResultResource` object and generates the Go code to build this object using builders.
func SearchResultResourceConverter(input SearchResultResource) string {
	calls := []string{
		`alertingv0alpha1.NewSearchResultResourceBuilder()`,
	}
	var buffer strings.Builder
	if input.Group != "" {

		buffer.WriteString(`Group(`)
		arg0 := fmt.Sprintf("%#v", input.Group)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Resource != "" {

		buffer.WriteString(`Resource(`)
		arg0 := fmt.Sprintf("%#v", input.Resource)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Kind != "" {

		buffer.WriteString(`Kind(`)
		arg0 := fmt.Sprintf("%#v", input.Kind)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Name != "" {

		buffer.WriteString(`Name(`)
		arg0 := fmt.Sprintf("%#v", input.Name)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
