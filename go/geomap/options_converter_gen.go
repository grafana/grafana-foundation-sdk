// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package geomap

import (
	"strings"

	common "github.com/grafana/grafana-foundation-sdk/go/common"
)

// OptionsConverter accepts a `Options` object and generates the Go code to build this object using builders.
func OptionsConverter(input Options) string {
	calls := []string{
		`geomap.NewOptionsBuilder()`,
	}
	var buffer strings.Builder

	{
		buffer.WriteString(`View(`)
		arg0 := MapViewConfigConverter(input.View)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	{
		buffer.WriteString(`Controls(`)
		arg0 := ControlsOptionsConverter(input.Controls)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	{
		buffer.WriteString(`Basemap(`)
		arg0 := common.MapLayerOptionsConverter(input.Basemap)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.Layers != nil && len(input.Layers) >= 1 {

		buffer.WriteString(`Layers(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Layers {
			tmplayersarg1 := common.MapLayerOptionsConverter(arg1)
			tmparg0 = append(tmparg0, tmplayersarg1)
		}
		arg0 := "[]cog.Builder[common.MapLayerOptions]{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	{
		buffer.WriteString(`Tooltip(`)
		arg0 := TooltipOptionsConverter(input.Tooltip)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	return strings.Join(calls, ".\t\n")
}
