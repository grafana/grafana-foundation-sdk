// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	"fmt"
	"strings"
)

// FiltersConverter accepts a `Filters` object and generates the Go code to build this object using builders.
func FiltersConverter(input Filters) string {
	calls := []string{
		`elasticsearch.NewFiltersBuilder()`,
	}
	var buffer strings.Builder
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
		arg0 := ElasticsearchFiltersSettingsConverter(*input.Settings)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
