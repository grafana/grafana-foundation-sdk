// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// QueryVariableConverter accepts a `QueryVariable` object and generates the Go code to build this object using builders.
func QueryVariableConverter(input QueryVariableKind) string {
	calls := []string{
		`dashboardv2beta1.NewQueryVariableBuilder(` + fmt.Sprintf("%#v", input.Spec.Name) + `)`,
	}
	var buffer strings.Builder
	if input.Spec.Name != "" {

		buffer.WriteString(`Name(`)
		arg0 := fmt.Sprintf("%#v", input.Spec.Name)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	{
		buffer.WriteString(`Current(`)
		arg0 := cog.Dump(input.Spec.Current)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.Spec.Label != nil && *input.Spec.Label != "" {

		buffer.WriteString(`Label(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.Label)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	{
		buffer.WriteString(`Hide(`)
		arg0 := cog.Dump(input.Spec.Hide)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	{
		buffer.WriteString(`Refresh(`)
		arg0 := cog.Dump(input.Spec.Refresh)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.Spec.SkipUrlSync != false {

		buffer.WriteString(`SkipUrlSync(`)
		arg0 := fmt.Sprintf("%#v", input.Spec.SkipUrlSync)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.Description != nil && *input.Spec.Description != "" {

		buffer.WriteString(`Description(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.Description)
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

	if input.Spec.Regex != "" {

		buffer.WriteString(`Regex(`)
		arg0 := fmt.Sprintf("%#v", input.Spec.Regex)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	{
		buffer.WriteString(`Sort(`)
		arg0 := cog.Dump(input.Spec.Sort)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.Spec.Definition != nil && *input.Spec.Definition != "" {

		buffer.WriteString(`Definition(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.Definition)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.Options != nil && len(input.Spec.Options) >= 1 {

		buffer.WriteString(`Options(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Spec.Options {
			tmpoptionsarg1 := cog.Dump(arg1)
			tmparg0 = append(tmparg0, tmpoptionsarg1)
		}
		arg0 := "[]dashboardv2beta1.VariableOption{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.Multi != false {

		buffer.WriteString(`Multi(`)
		arg0 := fmt.Sprintf("%#v", input.Spec.Multi)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.IncludeAll != false {

		buffer.WriteString(`IncludeAll(`)
		arg0 := fmt.Sprintf("%#v", input.Spec.IncludeAll)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.AllValue != nil && *input.Spec.AllValue != "" {

		buffer.WriteString(`AllValue(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.AllValue)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.Placeholder != nil && *input.Spec.Placeholder != "" {

		buffer.WriteString(`Placeholder(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.Placeholder)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.AllowCustomValue != true {

		buffer.WriteString(`AllowCustomValue(`)
		arg0 := fmt.Sprintf("%#v", input.Spec.AllowCustomValue)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.StaticOptions != nil && len(input.Spec.StaticOptions) >= 1 {

		buffer.WriteString(`StaticOptions(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Spec.StaticOptions {
			tmpstaticOptionsarg1 := cog.Dump(arg1)
			tmparg0 = append(tmparg0, tmpstaticOptionsarg1)
		}
		arg0 := "[]dashboardv2beta1.VariableOption{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.StaticOptionsOrder != nil {

		buffer.WriteString(`StaticOptionsOrder(`)
		arg0 := cog.Dump(*input.Spec.StaticOptionsOrder)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
