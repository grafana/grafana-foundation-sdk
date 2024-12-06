// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// AzureMonitorQueryConverter accepts a `AzureMonitorQuery` object and generates the Go code to build this object using builders.
func AzureMonitorQueryConverter(input AzureMonitorQuery) string {
	calls := []string{
		`azuremonitor.NewAzureMonitorQueryBuilder()`,
	}
	var buffer strings.Builder
	if input.RefId != "" {

		buffer.WriteString(`RefId(`)
		arg0 := fmt.Sprintf("%#v", input.RefId)
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
	if input.QueryType != nil && *input.QueryType != "" {

		buffer.WriteString(`QueryType(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.QueryType))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Subscription != nil && *input.Subscription != "" {

		buffer.WriteString(`Subscription(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Subscription))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Subscriptions != nil && len(input.Subscriptions) >= 1 {

		buffer.WriteString(`Subscriptions(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Subscriptions {
			tmpsubscriptionsarg1 := fmt.Sprintf("%#v", arg1)
			tmparg0 = append(tmparg0, tmpsubscriptionsarg1)
		}
		arg0 := "[]string{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.AzureMonitor != nil {

		buffer.WriteString(`AzureMonitor(`)
		arg0 := AzureMetricQueryConverter(*input.AzureMonitor)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.AzureLogAnalytics != nil {

		buffer.WriteString(`AzureLogAnalytics(`)
		arg0 := AzureLogsQueryConverter(*input.AzureLogAnalytics)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.AzureResourceGraph != nil {

		buffer.WriteString(`AzureResourceGraph(`)
		arg0 := AzureResourceGraphQueryConverter(*input.AzureResourceGraph)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.AzureTraces != nil {

		buffer.WriteString(`AzureTraces(`)
		arg0 := AzureTracesQueryConverter(*input.AzureTraces)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.GrafanaTemplateVariableFn != nil {

		buffer.WriteString(`GrafanaTemplateVariableFn(`)
		arg0 := cog.Dump(*input.GrafanaTemplateVariableFn)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.ResourceGroup != nil && *input.ResourceGroup != "" {

		buffer.WriteString(`ResourceGroup(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.ResourceGroup))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Namespace != nil && *input.Namespace != "" {

		buffer.WriteString(`Namespace(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Namespace))
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
	if input.Region != nil && *input.Region != "" {

		buffer.WriteString(`Region(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Region))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Datasource != nil {

		buffer.WriteString(`Datasource(`)
		arg0 := cog.Dump(*input.Datasource)
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
