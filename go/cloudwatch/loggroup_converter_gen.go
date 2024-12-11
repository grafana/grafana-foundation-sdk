// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package cloudwatch

import (
	"fmt"
	"strings"
)

// LogGroupConverter accepts a `LogGroup` object and generates the Go code to build this object using builders.
func LogGroupConverter(input LogGroup) string {
	calls := []string{
		`cloudwatch.NewLogGroupBuilder()`,
	}
	var buffer strings.Builder
	if input.Arn != "" {

		buffer.WriteString(`Arn(`)
		arg0 := fmt.Sprintf("%#v", input.Arn)
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
	if input.AccountId != nil && *input.AccountId != "" {

		buffer.WriteString(`AccountId(`)
		arg0 := fmt.Sprintf("%#v", *input.AccountId)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.AccountLabel != nil && *input.AccountLabel != "" {

		buffer.WriteString(`AccountLabel(`)
		arg0 := fmt.Sprintf("%#v", *input.AccountLabel)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
