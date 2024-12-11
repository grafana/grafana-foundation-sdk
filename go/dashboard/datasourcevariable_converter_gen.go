// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboard

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// DatasourceVariableConverter accepts a `DatasourceVariable` object and generates the Go code to build this object using builders.
func DatasourceVariableConverter(input VariableModel) string {
	calls := []string{
		`dashboard.NewDatasourceVariableBuilder(` + fmt.Sprintf("%#v", input.Name) + `)`,
	}
	var buffer strings.Builder
	if input.Id != "" && input.Id != "00000000-0000-0000-0000-000000000000" {

		buffer.WriteString(`Id(`)
		arg0 := fmt.Sprintf("%#v", input.Id)
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
	if input.Label != nil && *input.Label != "" {

		buffer.WriteString(`Label(`)
		arg0 := fmt.Sprintf("%#v", *input.Label)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	{
		buffer.WriteString(`Hide(`)
		arg0 := cog.Dump(input.Hide)
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
	if input.Query != nil && input.Query.String != nil && *input.Query.String != "" {

		buffer.WriteString(`Type(`)
		arg0 := fmt.Sprintf("%#v", *input.Query.String)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.AllFormat != nil && *input.AllFormat != "" {

		buffer.WriteString(`AllFormat(`)
		arg0 := fmt.Sprintf("%#v", *input.AllFormat)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Current != nil {

		buffer.WriteString(`Current(`)
		arg0 := cog.Dump(*input.Current)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Multi != nil && *input.Multi != false {

		buffer.WriteString(`Multi(`)
		arg0 := fmt.Sprintf("%#v", *input.Multi)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.IncludeAll != nil && *input.IncludeAll != false {

		buffer.WriteString(`IncludeAll(`)
		arg0 := fmt.Sprintf("%#v", *input.IncludeAll)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.AllValue != nil && *input.AllValue != "" {

		buffer.WriteString(`AllValue(`)
		arg0 := fmt.Sprintf("%#v", *input.AllValue)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Regex != nil && *input.Regex != "" {

		buffer.WriteString(`Regex(`)
		arg0 := fmt.Sprintf("%#v", *input.Regex)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
