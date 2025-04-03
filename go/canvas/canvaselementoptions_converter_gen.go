// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package canvas

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// CanvasElementOptionsConverter accepts a `CanvasElementOptions` object and generates the Go code to build this object using builders.
func CanvasElementOptionsConverter(input CanvasElementOptions) string {
	calls := []string{
		`canvas.NewCanvasElementOptionsBuilder()`,
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
	if input.Type != "" {

		buffer.WriteString(`Type(`)
		arg0 := fmt.Sprintf("%#v", input.Type)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Config != nil {

		buffer.WriteString(`Config(`)
		arg0 := cog.Dump(input.Config)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Constraint != nil {

		buffer.WriteString(`Constraint(`)
		arg0 := ConstraintConverter(*input.Constraint)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Placement != nil {

		buffer.WriteString(`Placement(`)
		arg0 := PlacementConverter(*input.Placement)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Background != nil {

		buffer.WriteString(`Background(`)
		arg0 := BackgroundConfigConverter(*input.Background)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Border != nil {

		buffer.WriteString(`Border(`)
		arg0 := LineConfigConverter(*input.Border)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Connections != nil && len(input.Connections) >= 1 {

		buffer.WriteString(`Connections(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Connections {
			tmpconnectionsarg1 := CanvasConnectionConverter(arg1)
			tmparg0 = append(tmparg0, tmpconnectionsarg1)
		}
		arg0 := "[]cog.Builder[canvas.CanvasConnection]{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
