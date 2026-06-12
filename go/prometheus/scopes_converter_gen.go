// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package prometheus

import (
	"fmt"
	"strings"
)

// ScopesConverter accepts a `Scopes` object and generates the Go code to build this object using builders.
func ScopesConverter(input Scopes) string {
	calls := []string{
		`prometheus.NewScopesBuilder()`,
	}
	var buffer strings.Builder
	if input.DefaultPath != nil && len(input.DefaultPath) >= 1 {

		buffer.WriteString(`DefaultPath(`)
		tmparg0 := []string{}
		for _, arg1 := range input.DefaultPath {
			tmpdefaultPatharg1 := fmt.Sprintf("%#v", arg1)
			tmparg0 = append(tmparg0, tmpdefaultPatharg1)
		}
		arg0 := "[]string{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Filters != nil && len(input.Filters) >= 1 {

		buffer.WriteString(`Filters(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Filters {
			tmpfiltersarg1 := ScopesFiltersConverter(arg1)
			tmparg0 = append(tmparg0, tmpfiltersarg1)
		}
		arg0 := "[]cog.Builder[prometheus.ScopesFilters]{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Title != "" {

		buffer.WriteString(`Title(`)
		arg0 := fmt.Sprintf("%#v", input.Title)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
