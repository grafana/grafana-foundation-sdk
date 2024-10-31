// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// ExtendedStatsConverter accepts a `ExtendedStats` object and generates the Go code to build this object using builders.
func ExtendedStatsConverter(input ExtendedStats) string {
	calls := []string{
		`elasticsearch.NewExtendedStatsBuilder()`,
	}
	var buffer strings.Builder
	if input.Settings != nil {

		buffer.WriteString(`Settings(`)
		arg0 := ElasticsearchExtendedStatsSettingsConverter(*input.Settings)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
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
	if input.Meta != nil {

		buffer.WriteString(`Meta(`)
		arg0 := cog.Dump(input.Meta)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Hide != nil {

		buffer.WriteString(`Hide(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Hide))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
