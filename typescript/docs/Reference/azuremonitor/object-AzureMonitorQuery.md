---
title: <span class="badge object-type-interface"></span> AzureMonitorQuery
---
# <span class="badge object-type-interface"></span> AzureMonitorQuery

## Definition

```typescript
export interface AzureMonitorQuery {
	// A unique identifier for the query within the list of targets.
	// In server side expressions, the refId is used as a variable name to identify results.
	// By default, the UI will assign A->Z; however setting meaningful names may be useful.
	refId: string;
	// If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
	hide?: boolean;
	// Specify the query flavor
	// TODO make this required and give it a default
	queryType?: string;
	// Azure subscription containing the resource(s) to be queried.
	subscription?: string;
	// Subscriptions to be queried via Azure Resource Graph.
	subscriptions?: string[];
	// Azure Monitor Metrics sub-query properties.
	azureMonitor?: azuremonitor.AzureMetricQuery;
	// Azure Monitor Logs sub-query properties.
	azureLogAnalytics?: azuremonitor.AzureLogsQuery;
	// Azure Resource Graph sub-query properties.
	azureResourceGraph?: azuremonitor.AzureResourceGraphQuery;
	// Application Insights Traces sub-query properties.
	azureTraces?: azuremonitor.AzureTracesQuery;
	// @deprecated Legacy template variable support.
	grafanaTemplateVariableFn?: azuremonitor.GrafanaTemplateVariableQuery;
	// Template variables params. These exist for backwards compatiblity with legacy template variables.
	resourceGroup?: string;
	namespace?: string;
	resource?: string;
	// For mixed data sources the selected datasource is on the query level.
	// For non mixed scenarios this is undefined.
	// TODO find a better way to do this ^ that's friendly to schema
	// TODO this shouldn't be unknown but DataSourceRef | null
	datasource?: dashboard.DataSourceRef;
	// Azure Monitor query type.
	// queryType: #AzureQueryType
	region?: string;
	_implementsDataqueryVariant(): void;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [AzureMonitorQueryBuilder](./builder-AzureMonitorQueryBuilder.md)
