// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// TargetConverter accepts a `Target` object and generates the Go code to build this object using builders.
func TargetConverter(input PanelQueryKind) string {
	calls := []string{
		`dashboardv2beta1.NewTargetBuilder()`,
	}
	var buffer strings.Builder

	{
		buffer.WriteString(`Query(`)
		arg0 := cog.ConvertDataQueryKindToCode(input.Spec.Query, input.Spec.Query.Group)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.Spec.RefId != "" {

		buffer.WriteString(`RefId(`)
		arg0 := fmt.Sprintf("%#v", input.Spec.RefId)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	{
		buffer.WriteString(`Hidden(`)
		arg0 := fmt.Sprintf("%#v", input.Spec.Hidden)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	return strings.Join(calls, ".\t\n")
}
