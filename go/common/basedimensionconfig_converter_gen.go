// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package common

import (
	"fmt"
	"strings"
)

// BaseDimensionConfigConverter accepts a `BaseDimensionConfig` object and generates the Go code to build this object using builders.
func BaseDimensionConfigConverter(input BaseDimensionConfig) string {
	calls := []string{
		`common.NewBaseDimensionConfigBuilder()`,
	}
	var buffer strings.Builder
	if input.Field != nil && *input.Field != "" {

		buffer.WriteString(`Field(`)
		arg0 := fmt.Sprintf("%#v", *input.Field)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
