// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// GeoHashGridConverter accepts a `GeoHashGrid` object and generates the Go code to build this object using builders.
func GeoHashGridConverter(input GeoHashGrid) string {
	calls := []string{
		`elasticsearch.NewGeoHashGridBuilder()`,
	}
	var buffer strings.Builder
	if input.Field != nil && *input.Field != "" {

		buffer.WriteString(`Field(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Field))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Id != "" {

		buffer.WriteString(`Id(`)
		arg0 := fmt.Sprintf("%#v", input.Id)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Settings != nil {

		buffer.WriteString(`Settings(`)
		arg0 := ElasticsearchGeoHashGridSettingsConverter(*input.Settings)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
