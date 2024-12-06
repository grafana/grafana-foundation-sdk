// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// AzureTracesQueryConverter accepts a `AzureTracesQuery` object and generates the Go code to build this object using builders.
func AzureTracesQueryConverter(input AzureTracesQuery) string {
	calls := []string{
		`azuremonitor.NewAzureTracesQueryBuilder()`,
	}
	var buffer strings.Builder
	if input.ResultFormat != nil {

		buffer.WriteString(`ResultFormat(`)
		arg0 := cog.Dump(*input.ResultFormat)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Resources != nil && len(input.Resources) >= 1 {

		buffer.WriteString(`Resources(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Resources {
			tmpresourcesarg1 := fmt.Sprintf("%#v", arg1)
			tmparg0 = append(tmparg0, tmpresourcesarg1)
		}
		arg0 := "[]string{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.OperationId != nil && *input.OperationId != "" {

		buffer.WriteString(`OperationId(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.OperationId))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.TraceTypes != nil && len(input.TraceTypes) >= 1 {

		buffer.WriteString(`TraceTypes(`)
		tmparg0 := []string{}
		for _, arg1 := range input.TraceTypes {
			tmptraceTypesarg1 := fmt.Sprintf("%#v", arg1)
			tmparg0 = append(tmparg0, tmptraceTypesarg1)
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
			tmpfiltersarg1 := AzureTracesFilterConverter(arg1)
			tmparg0 = append(tmparg0, tmpfiltersarg1)
		}
		arg0 := "[]cog.Builder[azuremonitor.AzureTracesFilter]{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Query != nil && *input.Query != "" {

		buffer.WriteString(`Query(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Query))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
