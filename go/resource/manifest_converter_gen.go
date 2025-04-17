// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package resource

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// ManifestConverter accepts a `Manifest` object and generates the Go code to build this object using builders.
func ManifestConverter(input Manifest) string {
	calls := []string{
		`resource.NewManifestBuilder()`,
	}
	var buffer strings.Builder
	if input.ApiVersion != "" {

		buffer.WriteString(`ApiVersion(`)
		arg0 := fmt.Sprintf("%#v", input.ApiVersion)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Kind != "" {

		buffer.WriteString(`Kind(`)
		arg0 := fmt.Sprintf("%#v", input.Kind)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	{
		buffer.WriteString(`Metadata(`)
		arg0 := MetadataConverter(input.Metadata)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.Spec != nil {

		buffer.WriteString(`Spec(`)
		arg0 := cog.Dump(input.Spec)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
