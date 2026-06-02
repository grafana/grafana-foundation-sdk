// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboard

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// AnnotationEventFieldMappingConverter accepts a `AnnotationEventFieldMapping` object and generates the Go code to build this object using builders.
func AnnotationEventFieldMappingConverter(input AnnotationEventFieldMapping) string {
	calls := []string{
		`dashboard.NewAnnotationEventFieldMappingBuilder()`,
	}
	var buffer strings.Builder
	if input.Source != nil {

		buffer.WriteString(`Source(`)
		arg0 := cog.Dump(*input.Source)
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
