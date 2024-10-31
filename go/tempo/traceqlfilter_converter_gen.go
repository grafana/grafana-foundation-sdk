// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package tempo

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// TraceqlFilterConverter accepts a `TraceqlFilter` object and generates the Go code to build this object using builders.
func TraceqlFilterConverter(input TraceqlFilter) string {
	calls := []string{
		`tempo.NewTraceqlFilterBuilder()`,
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
	if input.Tag != nil && *input.Tag != "" {

		buffer.WriteString(`Tag(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Tag))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Operator != nil && *input.Operator != "" {

		buffer.WriteString(`Operator(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Operator))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Value != nil && input.Value.ArrayOfString != nil && len(input.Value.ArrayOfString) >= 1 {

		buffer.WriteString(`Value(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Value.ArrayOfString {
			tmpArrayOfStringarg1 := fmt.Sprintf("%#v", arg1)
			tmparg0 = append(tmparg0, tmpArrayOfStringarg1)
		}
		arg0 := "[]string{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.ValueType != nil && *input.ValueType != "" {

		buffer.WriteString(`ValueType(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.ValueType))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Scope != nil {

		buffer.WriteString(`Scope(`)
		arg0 := cog.Dump(*input.Scope)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
