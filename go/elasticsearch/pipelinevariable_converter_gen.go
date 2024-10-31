// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	"fmt"
	"strings"
)

// PipelineVariableConverter accepts a `PipelineVariable` object and generates the Go code to build this object using builders.
func PipelineVariableConverter(input PipelineVariable) string {
	calls := []string{
		`elasticsearch.NewPipelineVariableBuilder()`,
	}
	var buffer strings.Builder
	if input.Name != "" {

		buffer.WriteString(`Name(`)
		arg0 := fmt.Sprintf("%#v", input.Name)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.PipelineAgg != "" {

		buffer.WriteString(`PipelineAgg(`)
		arg0 := fmt.Sprintf("%#v", input.PipelineAgg)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
