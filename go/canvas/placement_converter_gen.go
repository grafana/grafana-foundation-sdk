// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package canvas

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// PlacementConverter accepts a `Placement` object and generates the Go code to build this object using builders.
func PlacementConverter(input Placement) string {
	calls := []string{
		`canvas.NewPlacementBuilder()`,
	}
	var buffer strings.Builder
	if input.Top != nil {

		buffer.WriteString(`Top(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Top))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Left != nil {

		buffer.WriteString(`Left(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Left))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Right != nil {

		buffer.WriteString(`Right(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Right))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Bottom != nil {

		buffer.WriteString(`Bottom(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Bottom))
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
	if input.Height != nil {

		buffer.WriteString(`Height(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Height))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
