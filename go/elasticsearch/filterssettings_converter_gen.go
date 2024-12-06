// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	"strings"
)

// FiltersSettingsConverter accepts a `FiltersSettings` object and generates the Go code to build this object using builders.
func FiltersSettingsConverter(input FiltersSettings) string {
	calls := []string{
		`elasticsearch.NewFiltersSettingsBuilder()`,
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
