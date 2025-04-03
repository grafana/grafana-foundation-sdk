// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package athena

import (
	"fmt"
	"strings"
)

// ConnectionArgsConverter accepts a `ConnectionArgs` object and generates the Go code to build this object using builders.
func ConnectionArgsConverter(input ConnectionArgs) string {
	calls := []string{
		`athena.NewConnectionArgsBuilder()`,
	}
	var buffer strings.Builder
	if input.Region != nil && *input.Region != "" && *input.Region != "__default" {

		buffer.WriteString(`Region(`)
		arg0 := fmt.Sprintf("%#v", *input.Region)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Catalog != nil && *input.Catalog != "" && *input.Catalog != "__default" {

		buffer.WriteString(`Catalog(`)
		arg0 := fmt.Sprintf("%#v", *input.Catalog)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Database != nil && *input.Database != "" && *input.Database != "__default" {

		buffer.WriteString(`Database(`)
		arg0 := fmt.Sprintf("%#v", *input.Database)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.ResultReuseEnabled != nil && *input.ResultReuseEnabled != false {

		buffer.WriteString(`ResultReuseEnabled(`)
		arg0 := fmt.Sprintf("%#v", *input.ResultReuseEnabled)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.ResultReuseMaxAgeInMinutes != nil && *input.ResultReuseMaxAgeInMinutes != 60 {

		buffer.WriteString(`ResultReuseMaxAgeInMinutes(`)
		arg0 := fmt.Sprintf("%#v", *input.ResultReuseMaxAgeInMinutes)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
