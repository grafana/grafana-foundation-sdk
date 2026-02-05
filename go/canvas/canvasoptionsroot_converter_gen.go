// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package canvas

import (
	"fmt"
	"strings"
)

// CanvasOptionsRootConverter accepts a `CanvasOptionsRoot` object and generates the Go code to build this object using builders.
func CanvasOptionsRootConverter(input CanvasOptionsRoot) string {
	calls := []string{
		`canvas.NewCanvasOptionsRootBuilder()`,
	}
	var buffer strings.Builder
	if input.Name != "" {

		buffer.WriteString(`Name(`)
		arg0 := fmt.Sprintf("%#v", input.Name)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Elements != nil && len(input.Elements) >= 1 {

		buffer.WriteString(`Elements(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Elements {
			tmpelementsarg1 := CanvasElementOptionsConverter(arg1)
			tmparg0 = append(tmparg0, tmpelementsarg1)
		}
		arg0 := "[]cog.Builder[canvas.CanvasElementOptions]{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
