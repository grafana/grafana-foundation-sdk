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
	if input.Labels != nil {

		buffer.WriteString(`Labels(`)
		arg0 := "map[string]string{"
		for key, arg1 := range input.Labels {
			tmplabelsarg1 := fmt.Sprintf("%#v", arg1)
			arg0 += "\t" + fmt.Sprintf("%#v", key) + ": " + tmplabelsarg1 + ","
		}
		arg0 += "}"
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
	if input.Spec != nil && input.Spec.(*MonitorQuery).RefId != nil && *input.Spec.(*MonitorQuery).RefId != "" {

		buffer.WriteString(`RefId(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*MonitorQuery).RefId)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*MonitorQuery).Hide != nil {

		buffer.WriteString(`Hide(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*MonitorQuery).Hide)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*MonitorQuery).QueryType != nil && *input.Spec.(*MonitorQuery).QueryType != "" {

		buffer.WriteString(`QueryType(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*MonitorQuery).QueryType)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*MonitorQuery).Subscription != nil && *input.Spec.(*MonitorQuery).Subscription != "" {

		buffer.WriteString(`Subscription(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*MonitorQuery).Subscription)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*MonitorQuery).Subscriptions != nil && len(input.Spec.(*MonitorQuery).Subscriptions) >= 1 {

		buffer.WriteString(`Subscriptions(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Spec.(*MonitorQuery).Subscriptions {
			tmpsubscriptionsarg1 := fmt.Sprintf("%#v", arg1)
			tmparg0 = append(tmparg0, tmpsubscriptionsarg1)
		}
		arg0 := "[]string{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*MonitorQuery).AzureMonitor != nil {

		buffer.WriteString(`AzureMonitor(`)
		arg0 := MetricQueryConverter(*input.Spec.(*MonitorQuery).AzureMonitor)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*MonitorQuery).AzureLogAnalytics != nil {

		buffer.WriteString(`AzureLogAnalytics(`)
		arg0 := LogsQueryConverter(*input.Spec.(*MonitorQuery).AzureLogAnalytics)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*MonitorQuery).AzureResourceGraph != nil {

		buffer.WriteString(`AzureResourceGraph(`)
		arg0 := ResourceGraphQueryConverter(*input.Spec.(*MonitorQuery).AzureResourceGraph)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*MonitorQuery).AzureTraces != nil {

		buffer.WriteString(`AzureTraces(`)
		arg0 := TracesQueryConverter(*input.Spec.(*MonitorQuery).AzureTraces)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*MonitorQuery).GrafanaTemplateVariableFn != nil {

		buffer.WriteString(`GrafanaTemplateVariableFn(`)
		arg0 := cog.Dump(*input.Spec.(*MonitorQuery).GrafanaTemplateVariableFn)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*MonitorQuery).ResourceGroup != nil && *input.Spec.(*MonitorQuery).ResourceGroup != "" {

		buffer.WriteString(`ResourceGroup(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*MonitorQuery).ResourceGroup)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*MonitorQuery).Namespace != nil && *input.Spec.(*MonitorQuery).Namespace != "" {

		buffer.WriteString(`Namespace(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*MonitorQuery).Namespace)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*MonitorQuery).Resource != nil && *input.Spec.(*MonitorQuery).Resource != "" {

		buffer.WriteString(`Resource(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*MonitorQuery).Resource)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*MonitorQuery).Region != nil && *input.Spec.(*MonitorQuery).Region != "" {

		buffer.WriteString(`Region(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*MonitorQuery).Region)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*MonitorQuery).CustomNamespace != nil && *input.Spec.(*MonitorQuery).CustomNamespace != "" {

		buffer.WriteString(`CustomNamespace(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*MonitorQuery).CustomNamespace)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*MonitorQuery).Query != nil && *input.Spec.(*MonitorQuery).Query != "" {

		buffer.WriteString(`Query(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*MonitorQuery).Query)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
