---
title: <span class="badge object-type-enum"></span> QueryType
---
# <span class="badge object-type-enum"></span> QueryType

Defines the supported queryTypes. GrafanaTemplateVariableFn is deprecated

## Definition

```go
type QueryType string
const (
	QueryTypeMonitor QueryType = "Azure Monitor"
	QueryTypeLogAnalytics QueryType = "Azure Log Analytics"
	QueryTypeResourceGraph QueryType = "Azure Resource Graph"
	QueryTypeTraces QueryType = "Azure Traces"
	QueryTypeSubscriptionsQuery QueryType = "Azure Subscriptions"
	QueryTypeResourceGroupsQuery QueryType = "Azure Resource Groups"
	QueryTypeNamespacesQuery QueryType = "Azure Namespaces"
	QueryTypeResourceNamesQuery QueryType = "Azure Resource Names"
	QueryTypeMetricNamesQuery QueryType = "Azure Metric Names"
	QueryTypeWorkspacesQuery QueryType = "Azure Workspaces"
	QueryTypeLocationsQuery QueryType = "Azure Regions"
	QueryTypeGrafanaTemplateVariableFn QueryType = "Grafana Template Variable Function"
	QueryTypeTraceExemplar QueryType = "traceql"
	QueryTypeCustomNamespacesQuery QueryType = "Azure Custom Namespaces"
	QueryTypeCustomMetricNamesQuery QueryType = "Azure Custom Metric Names"
)

```
