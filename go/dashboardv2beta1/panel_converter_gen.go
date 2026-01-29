// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// PanelConverter accepts a `Panel` object and generates the Go code to build this object using builders.
func PanelConverter(input PanelKind) string {
	calls := []string{
		`dashboardv2beta1.NewPanelBuilder()`,
	}
	var buffer strings.Builder

	{
		buffer.WriteString(`Id(`)
		arg0 := fmt.Sprintf("%#v", input.Spec.Id)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.Spec.Title != "" {

		buffer.WriteString(`Title(`)
		arg0 := fmt.Sprintf("%#v", input.Spec.Title)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.Description != "" {

		buffer.WriteString(`Description(`)
		arg0 := fmt.Sprintf("%#v", input.Spec.Description)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.Links != nil && len(input.Spec.Links) >= 1 {

		buffer.WriteString(`Links(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Spec.Links {
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
		arg0 := QueryGroupConverter(input.Spec.Data)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	{
		buffer.WriteString(`Visualization(`)
		arg0 := cog.ConvertPanelToCode(input.Spec.VizConfig, input.Spec.VizConfig.Group)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.Spec.Transparent != nil {

		buffer.WriteString(`Transparent(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.Transparent)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
