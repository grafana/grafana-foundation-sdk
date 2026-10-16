// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// V2PanelSpecConverter accepts a `V2PanelSpec` object and generates the Go code to build this object using builders.
func V2PanelSpecConverter(input V2PanelSpec) string {
	calls := []string{
		`dashboardv2beta1.NewV2PanelSpecBuilder()`,
	}
	var buffer strings.Builder

	{
		buffer.WriteString(`Id(`)
		arg0 := fmt.Sprintf("%#v", input.Id)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.Title != "" {

		buffer.WriteString(`Title(`)
		arg0 := fmt.Sprintf("%#v", input.Title)
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
	if input.Subtitle != nil && *input.Subtitle != "" {

		buffer.WriteString(`Subtitle(`)
		arg0 := fmt.Sprintf("%#v", *input.Subtitle)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Links != nil && len(input.Links) >= 1 {

		buffer.WriteString(`Links(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Links {
			tmplinksarg1 := DataLinkConverter(arg1)
			tmparg0 = append(tmparg0, tmplinksarg1)
		}
		arg0 := "[]cog.Builder[dashboardv2beta1.DataLink]{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	{
		buffer.WriteString(`Data(`)
		arg0 := V2QueryGroupKindConverter(input.Data)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	{
		buffer.WriteString(`VizConfig(`)
		arg0 := cog.ConvertPanelToCode(input.VizConfig, input.VizConfig.Group)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.Transparent != nil {

		buffer.WriteString(`Transparent(`)
		arg0 := fmt.Sprintf("%#v", *input.Transparent)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
