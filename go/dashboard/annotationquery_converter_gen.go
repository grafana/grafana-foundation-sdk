// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboard

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// AnnotationQueryConverter accepts a `AnnotationQuery` object and generates the Go code to build this object using builders.
func AnnotationQueryConverter(input AnnotationQuery) string {
	calls := []string{
		`dashboard.NewAnnotationQueryBuilder()`,
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
	if input.Datasource != nil {

		buffer.WriteString(`Datasource(`)
		arg0 := cog.Dump(*input.Datasource)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Enable != true {

		buffer.WriteString(`Enable(`)
		arg0 := fmt.Sprintf("%#v", input.Enable)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Hide != nil && *input.Hide != false {

		buffer.WriteString(`Hide(`)
		arg0 := fmt.Sprintf("%#v", *input.Hide)
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
	if input.Filter != nil {

		buffer.WriteString(`Filter(`)
		arg0 := AnnotationPanelFilterConverter(*input.Filter)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Target != nil {

		buffer.WriteString(`Target(`)
		arg0 := cog.ConvertDataqueryToCode(input.Target)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Type != nil && *input.Type != "" {

		buffer.WriteString(`Type(`)
		arg0 := fmt.Sprintf("%#v", *input.Type)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.BuiltIn != nil && *input.BuiltIn != 0 {

		buffer.WriteString(`BuiltIn(`)
		arg0 := fmt.Sprintf("%#v", *input.BuiltIn)
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
	if input.Expr != nil && *input.Expr != "" {

		buffer.WriteString(`Expr(`)
		arg0 := fmt.Sprintf("%#v", *input.Expr)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.TextFormat != nil && *input.TextFormat != "" {

		buffer.WriteString(`TextFormat(`)
		arg0 := fmt.Sprintf("%#v", *input.TextFormat)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.TitleFormat != nil && *input.TitleFormat != "" {

		buffer.WriteString(`TitleFormat(`)
		arg0 := fmt.Sprintf("%#v", *input.TitleFormat)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.TagKeys != nil && *input.TagKeys != "" {

		buffer.WriteString(`TagKeys(`)
		arg0 := fmt.Sprintf("%#v", *input.TagKeys)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Step != nil && *input.Step != "" {

		buffer.WriteString(`Step(`)
		arg0 := fmt.Sprintf("%#v", *input.Step)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.UseValueForTime != nil {

		buffer.WriteString(`UseValueForTime(`)
		arg0 := fmt.Sprintf("%#v", *input.UseValueForTime)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Mappings != nil {

		buffer.WriteString(`Mappings(`)
		arg0 := "map[string]cog.Builder[dashboard.AnnotationEventFieldMapping]{"
		for key, arg1 := range input.Mappings {
			tmpmappingsarg1 := AnnotationEventFieldMappingConverter(arg1)
			arg0 += "\t" + fmt.Sprintf("%#v", key) + ": " + tmpmappingsarg1 + ","
		}
		arg0 += "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
