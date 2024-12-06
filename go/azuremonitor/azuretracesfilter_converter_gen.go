// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	"fmt"
	"strings"
)

// AzureTracesFilterConverter accepts a `AzureTracesFilter` object and generates the Go code to build this object using builders.
func AzureTracesFilterConverter(input AzureTracesFilter) string {
	calls := []string{
		`azuremonitor.NewAzureTracesFilterBuilder()`,
	}
	var buffer strings.Builder
	if input.Property != "" {

		buffer.WriteString(`Property(`)
		arg0 := fmt.Sprintf("%#v", input.Property)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Operation != "" {

		buffer.WriteString(`Operation(`)
		arg0 := fmt.Sprintf("%#v", input.Operation)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Filters != nil && len(input.Filters) >= 1 {

		buffer.WriteString(`Filters(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Filters {
			tmpfiltersarg1 := fmt.Sprintf("%#v", arg1)
			tmparg0 = append(tmparg0, tmpfiltersarg1)
		}
		arg0 := "[]string{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
