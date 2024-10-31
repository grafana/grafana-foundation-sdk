---
title: <span class="badge object-type-enum"></span> AzureQueryType
---
# <span class="badge object-type-enum"></span> AzureQueryType

Defines the supported queryTypes. GrafanaTemplateVariableFn is deprecated

## Definition

```go
type AzureQueryType string
const (
	AzureQueryTypeAzureMonitor AzureQueryType = "Azure Monitor"
	AzureQueryTypeLogAnalytics AzureQueryType = "Azure Log Analytics"
	AzureQueryTypeAzureResourceGraph AzureQueryType = "Azure Resource Graph"
	AzureQueryTypeAzureTraces AzureQueryType = "Azure Traces"
	AzureQueryTypeSubscriptionsQuery AzureQueryType = "Azure Subscriptions"
	AzureQueryTypeResourceGroupsQuery AzureQueryType = "Azure Resource Groups"
	AzureQueryTypeNamespacesQuery AzureQueryType = "Azure Namespaces"
	AzureQueryTypeResourceNamesQuery AzureQueryType = "Azure Resource Names"
	AzureQueryTypeMetricNamesQuery AzureQueryType = "Azure Metric Names"
	AzureQueryTypeWorkspacesQuery AzureQueryType = "Azure Workspaces"
	AzureQueryTypeLocationsQuery AzureQueryType = "Azure Regions"
	AzureQueryTypeGrafanaTemplateVariableFn AzureQueryType = "Grafana Template Variable Function"
	AzureQueryTypeTraceExemplar AzureQueryType = "traceql"
)

```
