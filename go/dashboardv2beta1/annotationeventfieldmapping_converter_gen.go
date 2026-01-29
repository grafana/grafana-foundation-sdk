// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"fmt"
	"strings"
)

// AnnotationEventFieldMappingConverter accepts a `AnnotationEventFieldMapping` object and generates the Go code to build this object using builders.
func AnnotationEventFieldMappingConverter(input AnnotationEventFieldMapping) string {
	calls := []string{
		`dashboardv2beta1.NewAnnotationEventFieldMappingBuilder()`,
	}
	var buffer strings.Builder
	if input.Source != nil && *input.Source != "" && *input.Source != "field" {

		buffer.WriteString(`Source(`)
		arg0 := fmt.Sprintf("%#v", *input.Source)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Value != nil && *input.Value != "" {

		buffer.WriteString(`Value(`)
		arg0 := fmt.Sprintf("%#v", *input.Value)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Regex != nil && *input.Regex != "" {

		buffer.WriteString(`Regex(`)
		arg0 := fmt.Sprintf("%#v", *input.Regex)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
