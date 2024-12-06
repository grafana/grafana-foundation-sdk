---
title: <span class="badge object-type-enum"></span> AzureQueryType
---
# <span class="badge object-type-enum"></span> AzureQueryType

Defines the supported queryTypes. GrafanaTemplateVariableFn is deprecated

## Definition

```typescript
export enum AzureQueryType {
	AzureMonitor = "Azure Monitor",
	LogAnalytics = "Azure Log Analytics",
	AzureResourceGraph = "Azure Resource Graph",
	AzureTraces = "Azure Traces",
	SubscriptionsQuery = "Azure Subscriptions",
	ResourceGroupsQuery = "Azure Resource Groups",
	NamespacesQuery = "Azure Namespaces",
	ResourceNamesQuery = "Azure Resource Names",
	MetricNamesQuery = "Azure Metric Names",
	WorkspacesQuery = "Azure Workspaces",
	LocationsQuery = "Azure Regions",
	GrafanaTemplateVariableFn = "Grafana Template Variable Function",
	TraceExemplar = "traceql",
}

```
