// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package canvas

import (
	"fmt"
	"strings"

	common "github.com/grafana/grafana-foundation-sdk/go/common"
)

// CanvasConnectionConverter accepts a `CanvasConnection` object and generates the Go code to build this object using builders.
func CanvasConnectionConverter(input CanvasConnection) string {
	calls := []string{
		`canvas.NewCanvasConnectionBuilder()`,
	}
	var buffer strings.Builder

	{
		buffer.WriteString(`Source(`)
		arg0 := ConnectionCoordinatesConverter(input.Source)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	{
		buffer.WriteString(`Target(`)
		arg0 := ConnectionCoordinatesConverter(input.Target)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.TargetName != nil && *input.TargetName != "" {

		buffer.WriteString(`TargetName(`)
		arg0 := fmt.Sprintf("%#v", *input.TargetName)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Color != nil {

		buffer.WriteString(`Color(`)
		arg0 := common.ColorDimensionConfigConverter(*input.Color)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Size != nil {

		buffer.WriteString(`Size(`)
		arg0 := common.ScaleDimensionConfigConverter(*input.Size)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
