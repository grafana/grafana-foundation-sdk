// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// TermsSettingsConverter accepts a `TermsSettings` object and generates the Go code to build this object using builders.
func TermsSettingsConverter(input TermsSettings) string {
	calls := []string{
		`elasticsearch.NewTermsSettingsBuilder()`,
	}
	var buffer strings.Builder
	if input.Order != nil {

		buffer.WriteString(`Order(`)
		arg0 := cog.Dump(*input.Order)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Size != nil && *input.Size != "" {

		buffer.WriteString(`Size(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Size))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.MinDocCount != nil && *input.MinDocCount != "" {

		buffer.WriteString(`MinDocCount(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.MinDocCount))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.OrderBy != nil && *input.OrderBy != "" {

		buffer.WriteString(`OrderBy(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.OrderBy))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Missing != nil && *input.Missing != "" {

		buffer.WriteString(`Missing(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Missing))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
