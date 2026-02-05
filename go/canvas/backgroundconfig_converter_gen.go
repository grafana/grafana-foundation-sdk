// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package canvas

import (
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
)

// BackgroundConfigConverter accepts a `BackgroundConfig` object and generates the Go code to build this object using builders.
func BackgroundConfigConverter(input BackgroundConfig) string {
	calls := []string{
		`canvas.NewBackgroundConfigBuilder()`,
	}
	var buffer strings.Builder
	if input.Color != nil {

		buffer.WriteString(`Color(`)
		arg0 := common.ColorDimensionConfigConverter(*input.Color)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Image != nil {

		buffer.WriteString(`Image(`)
		arg0 := common.ResourceDimensionConfigConverter(*input.Image)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Size != nil {

		buffer.WriteString(`Size(`)
		arg0 := cog.Dump(*input.Size)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
