// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package canvas

import (
	"fmt"
	"strings"
)

func ConnectionCoordinatesConverter(input ConnectionCoordinates) string {
	calls := []string{
		`canvas.NewConnectionCoordinatesBuilder()`,
	}
	var buffer strings.Builder

	{
		buffer.WriteString(`X(`)
		arg0 := fmt.Sprintf("%#v", input.X)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	{
		buffer.WriteString(`Y(`)
		arg0 := fmt.Sprintf("%#v", input.Y)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	return strings.Join(calls, ".\t\n")
}
