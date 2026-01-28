// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// AnnotationQueryConverter accepts a `AnnotationQuery` object and generates the Go code to build this object using builders.
func AnnotationQueryConverter(input AnnotationQueryKind) string {
	calls := []string{
		`dashboardv2beta1.NewAnnotationQueryBuilder()`,
	}
	var buffer strings.Builder

	{
		buffer.WriteString(`Spec(`)
		arg0 := AnnotationQuerySpecConverter(input.Spec)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	{
		buffer.WriteString(`Query(`)
		arg0 := cog.ConvertDataQueryKindToCode(input.Spec.Query, input.Spec.Query.Group)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	{
		buffer.WriteString(`Enable(`)
		arg0 := fmt.Sprintf("%#v", input.Spec.Enable)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	{
		buffer.WriteString(`Hide(`)
		arg0 := fmt.Sprintf("%#v", input.Spec.Hide)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.Spec.IconColor != "" {

		buffer.WriteString(`IconColor(`)
		arg0 := fmt.Sprintf("%#v", input.Spec.IconColor)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.Name != "" {

		buffer.WriteString(`Name(`)
		arg0 := fmt.Sprintf("%#v", input.Spec.Name)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.BuiltIn != nil && *input.Spec.BuiltIn != false {

		buffer.WriteString(`BuiltIn(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.BuiltIn)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.Filter != nil {

		buffer.WriteString(`Filter(`)
		arg0 := AnnotationPanelFilterConverter(*input.Spec.Filter)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.Placement != nil {

		buffer.WriteString(`Placement(`)
		arg0 := cog.Dump(*input.Spec.Placement)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.Mappings != nil {

		buffer.WriteString(`Mappings(`)
		arg0 := "map[string]cog.Builder[dashboardv2beta1.AnnotationEventFieldMapping]{"
		for key, arg1 := range input.Spec.Mappings {
			tmpmappingsarg1 := AnnotationEventFieldMappingConverter(arg1)
			arg0 += "\t" + fmt.Sprintf("%#v", key) + ": " + tmpmappingsarg1 + ","
		}
		arg0 += "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.LegacyOptions != nil {

		buffer.WriteString(`LegacyOptions(`)
		arg0 := "map[string]any{"
		for key, arg1 := range input.Spec.LegacyOptions {
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
