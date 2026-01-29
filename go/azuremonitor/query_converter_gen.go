// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	dashboardv2beta1 "github.com/grafana/grafana-foundation-sdk/go/dashboardv2beta1"
)

// QueryConverter accepts a `Query` object and generates the Go code to build this object using builders.
func QueryConverter(input dashboardv2beta1.DataQueryKind) string {
	calls := []string{
		`azuremonitor.NewQueryBuilder()`,
	}
	var buffer strings.Builder
	if input.Version != "" && input.Version != "v0" {

		buffer.WriteString(`Version(`)
		arg0 := fmt.Sprintf("%#v", input.Version)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Datasource != nil {

		buffer.WriteString(`Datasource(`)
		arg0 := dashboardv2beta1.Dashboardv2beta1DataQueryKindDatasourceConverter(*input.Datasource)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*AzureMonitorQuery).RefId != nil && *input.Spec.(*AzureMonitorQuery).RefId != "" {

		buffer.WriteString(`RefId(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*AzureMonitorQuery).RefId)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*AzureMonitorQuery).Hide != nil {

		buffer.WriteString(`Hide(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*AzureMonitorQuery).Hide)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*AzureMonitorQuery).QueryType != nil && *input.Spec.(*AzureMonitorQuery).QueryType != "" {

		buffer.WriteString(`QueryType(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*AzureMonitorQuery).QueryType)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*AzureMonitorQuery).Subscription != nil && *input.Spec.(*AzureMonitorQuery).Subscription != "" {

		buffer.WriteString(`Subscription(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*AzureMonitorQuery).Subscription)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*AzureMonitorQuery).Subscriptions != nil && len(input.Spec.(*AzureMonitorQuery).Subscriptions) >= 1 {

		buffer.WriteString(`Subscriptions(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Spec.(*AzureMonitorQuery).Subscriptions {
			tmpsubscriptionsarg1 := fmt.Sprintf("%#v", arg1)
			tmparg0 = append(tmparg0, tmpsubscriptionsarg1)
		}
		arg0 := "[]string{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*AzureMonitorQuery).AzureMonitor != nil {

		buffer.WriteString(`AzureMonitor(`)
		arg0 := AzureMetricQueryConverter(*input.Spec.(*AzureMonitorQuery).AzureMonitor)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*AzureMonitorQuery).AzureLogAnalytics != nil {

		buffer.WriteString(`AzureLogAnalytics(`)
		arg0 := AzureLogsQueryConverter(*input.Spec.(*AzureMonitorQuery).AzureLogAnalytics)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*AzureMonitorQuery).AzureResourceGraph != nil {

		buffer.WriteString(`AzureResourceGraph(`)
		arg0 := AzureResourceGraphQueryConverter(*input.Spec.(*AzureMonitorQuery).AzureResourceGraph)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*AzureMonitorQuery).AzureTraces != nil {

		buffer.WriteString(`AzureTraces(`)
		arg0 := AzureTracesQueryConverter(*input.Spec.(*AzureMonitorQuery).AzureTraces)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*AzureMonitorQuery).GrafanaTemplateVariableFn != nil {

		buffer.WriteString(`GrafanaTemplateVariableFn(`)
		arg0 := cog.Dump(*input.Spec.(*AzureMonitorQuery).GrafanaTemplateVariableFn)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*AzureMonitorQuery).ResourceGroup != nil && *input.Spec.(*AzureMonitorQuery).ResourceGroup != "" {

		buffer.WriteString(`ResourceGroup(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*AzureMonitorQuery).ResourceGroup)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*AzureMonitorQuery).Namespace != nil && *input.Spec.(*AzureMonitorQuery).Namespace != "" {

		buffer.WriteString(`Namespace(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*AzureMonitorQuery).Namespace)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*AzureMonitorQuery).Resource != nil && *input.Spec.(*AzureMonitorQuery).Resource != "" {

		buffer.WriteString(`Resource(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*AzureMonitorQuery).Resource)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*AzureMonitorQuery).Region != nil && *input.Spec.(*AzureMonitorQuery).Region != "" {

		buffer.WriteString(`Region(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*AzureMonitorQuery).Region)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*AzureMonitorQuery).CustomNamespace != nil && *input.Spec.(*AzureMonitorQuery).CustomNamespace != "" {

		buffer.WriteString(`CustomNamespace(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*AzureMonitorQuery).CustomNamespace)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*AzureMonitorQuery).Datasource != nil {

		buffer.WriteString(`OldDatasource(`)
		arg0 := cog.Dump(*input.Spec.(*AzureMonitorQuery).Datasource)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*AzureMonitorQuery).Query != nil && *input.Spec.(*AzureMonitorQuery).Query != "" {

		buffer.WriteString(`Query(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*AzureMonitorQuery).Query)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
