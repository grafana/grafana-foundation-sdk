---
title: <span class="badge object-type-interface"></span> AzureLogsQuery
---
# <span class="badge object-type-interface"></span> AzureLogsQuery

Azure Monitor Logs sub-query properties

## Definition

```typescript
export interface AzureLogsQuery {
	// KQL query to be executed.
	query?: string;
	// Specifies the format results should be returned as.
	resultFormat?: azuremonitor.ResultFormat;
	// Array of resource URIs to be queried.
	resources?: string[];
	// If set to true the dashboard time range will be used as a filter for the query. Otherwise the query time ranges will be used. Defaults to false.
	dashboardTime?: boolean;
	// If dashboardTime is set to true this value dictates which column the time filter will be applied to. Defaults to the first tables timeSpan column, the first datetime column found, or TimeGenerated
	timeColumn?: string;
	// If set to true the query will be run as a basic logs query
	basicLogsQuery?: boolean;
	// Workspace ID. This was removed in Grafana 8, but remains for backwards compat.
	workspace?: string;
	// @deprecated Use resources instead
	resource?: string;
	// @deprecated Use dashboardTime instead
	intersectTime?: boolean;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [AzureLogsQueryBuilder](./builder-AzureLogsQueryBuilder.md)
