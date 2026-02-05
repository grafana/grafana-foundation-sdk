// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboard

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// RowConverter accepts a `Row` object and generates the Go code to build this object using builders.
func RowConverter(input RowPanel) string {
	calls := []string{
		`dashboard.NewRowBuilder(` + fmt.Sprintf("%#v", *input.Title) + `)`,
	}
	var buffer strings.Builder
	if input.Collapsed != false {

		buffer.WriteString(`Collapsed(`)
		arg0 := fmt.Sprintf("%#v", input.Collapsed)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Title != nil && *input.Title != "" {

		buffer.WriteString(`Title(`)
		arg0 := fmt.Sprintf("%#v", *input.Title)
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
	if input.GridPos != nil {

		buffer.WriteString(`GridPos(`)
		arg0 := cog.Dump(*input.GridPos)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	{
		buffer.WriteString(`Id(`)
		arg0 := fmt.Sprintf("%#v", input.Id)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.Panels != nil && len(input.Panels) >= 1 && input.Collapsed == true {

		buffer.WriteString(`WithPanel(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Panels {
			tmppanelsarg1 := cog.ConvertPanelToCode(arg1, arg1.Type)
			tmparg0 = append(tmparg0, tmppanelsarg1)
		}
		arg0 := "[]cog.Builder[dashboard.Panel]{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Repeat != nil && *input.Repeat != "" {

		buffer.WriteString(`Repeat(`)
		arg0 := fmt.Sprintf("%#v", *input.Repeat)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
