// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// KindConverter accepts a `Kind` object and generates the Go code to build this object using builders.
func KindConverter(input Kind) string {
	calls := []string{
		`dashboardv2beta1.NewKindBuilder()`,
	}
	var buffer strings.Builder
	if input.Kind != "" {

		buffer.WriteString(`Kind(`)
		arg0 := fmt.Sprintf("%#v", input.Kind)
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
	if input.Metadata != nil {

		buffer.WriteString(`Metadata(`)
		arg0 := cog.Dump(input.Metadata)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
