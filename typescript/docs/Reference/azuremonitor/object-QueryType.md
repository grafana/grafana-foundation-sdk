---
title: <span class="badge object-type-enum"></span> QueryType
---
# <span class="badge object-type-enum"></span> QueryType

Defines the supported queryTypes. GrafanaTemplateVariableFn is deprecated

## Definition

```typescript
export enum QueryType {
	Monitor = "Azure Monitor",
	LogAnalytics = "Azure Log Analytics",
	ResourceGraph = "Azure Resource Graph",
	Traces = "Azure Traces",
	SubscriptionsQuery = "Azure Subscriptions",
	ResourceGroupsQuery = "Azure Resource Groups",
	NamespacesQuery = "Azure Namespaces",
	ResourceNamesQuery = "Azure Resource Names",
	MetricNamesQuery = "Azure Metric Names",
	WorkspacesQuery = "Azure Workspaces",
	LocationsQuery = "Azure Regions",
	GrafanaTemplateVariableFn = "Grafana Template Variable Function",
	TraceExemplar = "traceql",
	CustomNamespacesQuery = "Azure Custom Namespaces",
	CustomMetricNamesQuery = "Azure Custom Metric Names",
}

```
