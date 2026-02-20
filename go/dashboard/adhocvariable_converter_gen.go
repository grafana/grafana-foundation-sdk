// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboard

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// AdHocVariableConverter accepts a `AdHocVariable` object and generates the Go code to build this object using builders.
func AdHocVariableConverter(input VariableModel) string {
	calls := []string{
		`dashboard.NewAdHocVariableBuilder(` + fmt.Sprintf("%#v", input.Name) + `)`,
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
	if input.Label != nil && *input.Label != "" {

		buffer.WriteString(`Label(`)
		arg0 := fmt.Sprintf("%#v", *input.Label)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Hide != nil {

		buffer.WriteString(`Hide(`)
		arg0 := cog.Dump(*input.Hide)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Description != nil && *input.Description != "" {

		buffer.WriteString(`Description(`)
		arg0 := fmt.Sprintf("%#v", *input.Description)
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
	if input.AllowCustomValue != nil && *input.AllowCustomValue != true {

		buffer.WriteString(`AllowCustomValue(`)
		arg0 := fmt.Sprintf("%#v", *input.AllowCustomValue)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.StaticOptions != nil && len(input.StaticOptions) >= 1 {

		buffer.WriteString(`StaticOptions(`)
		tmparg0 := []string{}
		for _, arg1 := range input.StaticOptions {
			tmpstaticOptionsarg1 := cog.Dump(arg1)
			tmparg0 = append(tmparg0, tmpstaticOptionsarg1)
		}
		arg0 := "[]dashboard.VariableOption{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.StaticOptionsOrder != nil {

		buffer.WriteString(`StaticOptionsOrder(`)
		arg0 := cog.Dump(*input.StaticOptionsOrder)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Definition != nil && *input.Definition != "" {

		buffer.WriteString(`Definition(`)
		arg0 := fmt.Sprintf("%#v", *input.Definition)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
