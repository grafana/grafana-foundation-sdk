// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// AnnotationQuerySpecConverter accepts a `AnnotationQuerySpec` object and generates the Go code to build this object using builders.
func AnnotationQuerySpecConverter(input AnnotationQuerySpec) string {
	calls := []string{
		`dashboardv2beta1.NewAnnotationQuerySpecBuilder()`,
	}
	var buffer strings.Builder

	{
		buffer.WriteString(`Query(`)
		arg0 := cog.ConvertDataQueryKindToCode(input.Query, input.Query.Group)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	{
		buffer.WriteString(`Enable(`)
		arg0 := fmt.Sprintf("%#v", input.Enable)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	{
		buffer.WriteString(`Hide(`)
		arg0 := fmt.Sprintf("%#v", input.Hide)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.IconColor != "" {

		buffer.WriteString(`IconColor(`)
		arg0 := fmt.Sprintf("%#v", input.IconColor)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Name != "" {

		buffer.WriteString(`Name(`)
		arg0 := fmt.Sprintf("%#v", input.Name)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.BuiltIn != nil && *input.BuiltIn != false {

		buffer.WriteString(`BuiltIn(`)
		arg0 := fmt.Sprintf("%#v", *input.BuiltIn)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Filter != nil {

		buffer.WriteString(`Filter(`)
		arg0 := AnnotationPanelFilterConverter(*input.Filter)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Placement != nil {

		buffer.WriteString(`Placement(`)
		arg0 := cog.Dump(*input.Placement)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.LegacyOptions != nil {

		buffer.WriteString(`LegacyOptions(`)
		arg0 := "map[string]any{"
		for key, arg1 := range input.LegacyOptions {
			tmplegacyOptionsarg1 := cog.Dump(arg1)
			arg0 += "\t" + fmt.Sprintf("%#v", key) + ": " + tmplegacyOptionsarg1 + ","
		}
		arg0 += "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
