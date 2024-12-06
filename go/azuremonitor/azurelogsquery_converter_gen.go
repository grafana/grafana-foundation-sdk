// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// AzureLogsQueryConverter accepts a `AzureLogsQuery` object and generates the Go code to build this object using builders.
func AzureLogsQueryConverter(input AzureLogsQuery) string {
	calls := []string{
		`azuremonitor.NewAzureLogsQueryBuilder()`,
	}
	var buffer strings.Builder
	if input.Query != nil && *input.Query != "" {

		buffer.WriteString(`Query(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Query))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
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
	if input.DashboardTime != nil {

		buffer.WriteString(`DashboardTime(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.DashboardTime))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.TimeColumn != nil && *input.TimeColumn != "" {

		buffer.WriteString(`TimeColumn(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.TimeColumn))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.BasicLogsQuery != nil {

		buffer.WriteString(`BasicLogsQuery(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.BasicLogsQuery))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Workspace != nil && *input.Workspace != "" {

		buffer.WriteString(`Workspace(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Workspace))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Resource != nil && *input.Resource != "" {

		buffer.WriteString(`Resource(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Resource))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.IntersectTime != nil {

		buffer.WriteString(`IntersectTime(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.IntersectTime))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
