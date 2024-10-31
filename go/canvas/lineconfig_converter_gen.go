// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package canvas

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
)

// LineConfigConverter accepts a `LineConfig` object and generates the Go code to build this object using builders.
func LineConfigConverter(input LineConfig) string {
	calls := []string{
		`canvas.NewLineConfigBuilder()`,
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
	if input.Width != nil {

		buffer.WriteString(`Width(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Width))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Radius != nil {

		buffer.WriteString(`Radius(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Radius))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
