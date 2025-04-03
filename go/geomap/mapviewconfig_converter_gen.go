// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package geomap

import (
	"fmt"
	"strings"
)

// MapViewConfigConverter accepts a `MapViewConfig` object and generates the Go code to build this object using builders.
func MapViewConfigConverter(input MapViewConfig) string {
	calls := []string{
		`geomap.NewMapViewConfigBuilder()`,
	}
	var buffer strings.Builder
	if input.Id != "" && input.Id != "zero" {

		buffer.WriteString(`Id(`)
		arg0 := fmt.Sprintf("%#v", input.Id)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Lat != nil && *input.Lat != 0 {

		buffer.WriteString(`Lat(`)
		arg0 := fmt.Sprintf("%#v", *input.Lat)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Lon != nil && *input.Lon != 0 {

		buffer.WriteString(`Lon(`)
		arg0 := fmt.Sprintf("%#v", *input.Lon)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Zoom != nil && *input.Zoom != 1 {

		buffer.WriteString(`Zoom(`)
		arg0 := fmt.Sprintf("%#v", *input.Zoom)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.MinZoom != nil {

		buffer.WriteString(`MinZoom(`)
		arg0 := fmt.Sprintf("%#v", *input.MinZoom)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.MaxZoom != nil {

		buffer.WriteString(`MaxZoom(`)
		arg0 := fmt.Sprintf("%#v", *input.MaxZoom)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Padding != nil {

		buffer.WriteString(`Padding(`)
		arg0 := fmt.Sprintf("%#v", *input.Padding)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.AllLayers != nil && *input.AllLayers != true {

		buffer.WriteString(`AllLayers(`)
		arg0 := fmt.Sprintf("%#v", *input.AllLayers)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.LastOnly != nil {

		buffer.WriteString(`LastOnly(`)
		arg0 := fmt.Sprintf("%#v", *input.LastOnly)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Layer != nil && *input.Layer != "" {

		buffer.WriteString(`Layer(`)
		arg0 := fmt.Sprintf("%#v", *input.Layer)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Shared != nil {

		buffer.WriteString(`Shared(`)
		arg0 := fmt.Sprintf("%#v", *input.Shared)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
