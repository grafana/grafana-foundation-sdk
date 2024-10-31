// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package common

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// ColorDimensionConfigConverter accepts a `ColorDimensionConfig` object and generates the Go code to build this object using builders.
func ColorDimensionConfigConverter(input ColorDimensionConfig) string {
	calls := []string{
		`common.NewColorDimensionConfigBuilder()`,
	}
	var buffer strings.Builder
	if input.Fixed != nil && *input.Fixed != "" {

		buffer.WriteString(`Fixed(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Fixed))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Field != nil && *input.Field != "" {

		buffer.WriteString(`Field(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Field))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
