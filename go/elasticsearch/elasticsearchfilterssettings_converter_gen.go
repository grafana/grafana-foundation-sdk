// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	"strings"
)

// ElasticsearchFiltersSettingsConverter accepts a `ElasticsearchFiltersSettings` object and generates the Go code to build this object using builders.
func ElasticsearchFiltersSettingsConverter(input ElasticsearchFiltersSettings) string {
	calls := []string{
		`elasticsearch.NewElasticsearchFiltersSettingsBuilder()`,
	}
	var buffer strings.Builder
	if input.Filters != nil && len(input.Filters) >= 1 {

		buffer.WriteString(`Filters(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Filters {
			tmpfiltersarg1 := FilterConverter(arg1)
			tmparg0 = append(tmparg0, tmpfiltersarg1)
		}
		arg0 := "[]cog.Builder[elasticsearch.Filter]{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
