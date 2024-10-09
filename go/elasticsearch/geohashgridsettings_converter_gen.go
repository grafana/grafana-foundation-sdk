// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

func GeoHashGridSettingsConverter(input GeoHashGridSettings) string {
	calls := []string{
		`elasticsearch.NewGeoHashGridSettingsBuilder()`,
	}
	var buffer strings.Builder
	if input.Precision != nil && *input.Precision != "" {

		buffer.WriteString(`Precision(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Precision))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}