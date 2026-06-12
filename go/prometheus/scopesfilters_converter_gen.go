// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package prometheus

import (
	"fmt"
	"strings"
)

// ScopesFiltersConverter accepts a `ScopesFilters` object and generates the Go code to build this object using builders.
func ScopesFiltersConverter(input ScopesFilters) string {
	calls := []string{
		`prometheus.NewScopesFiltersBuilder()`,
	}
	var buffer strings.Builder
	if input.Key != "" {

		buffer.WriteString(`Key(`)
		arg0 := fmt.Sprintf("%#v", input.Key)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Operator != "" {

		buffer.WriteString(`Operator(`)
		arg0 := fmt.Sprintf("%#v", input.Operator)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Value != "" {

		buffer.WriteString(`Value(`)
		arg0 := fmt.Sprintf("%#v", input.Value)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Values != nil && len(input.Values) >= 1 {

		buffer.WriteString(`Values(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Values {
			tmpvaluesarg1 := fmt.Sprintf("%#v", arg1)
			tmparg0 = append(tmparg0, tmpvaluesarg1)
		}
		arg0 := "[]string{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
