// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package xychart

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// XYDimensionConfigConverter accepts a `XYDimensionConfig` object and generates the Go code to build this object using builders.
func XYDimensionConfigConverter(input XYDimensionConfig) string {
	calls := []string{
		`xychart.NewXYDimensionConfigBuilder()`,
	}
	var buffer strings.Builder

	{
		buffer.WriteString(`Frame(`)
		arg0 := fmt.Sprintf("%#v", input.Frame)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.X != nil && *input.X != "" {

		buffer.WriteString(`X(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.X))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Exclude != nil && len(input.Exclude) >= 1 {

		buffer.WriteString(`Exclude(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Exclude {
			tmpexcludearg1 := fmt.Sprintf("%#v", arg1)
			tmparg0 = append(tmparg0, tmpexcludearg1)
		}
		arg0 := "[]string{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
