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
	// If set to true the intersection of time ranges specified in the query and Grafana will be used. Otherwise the query time ranges will be used. Defaults to false
	intersectTime?: boolean;
	// Workspace ID. This was removed in Grafana 8, but remains for backwards compat
	workspace?: string;
	// @deprecated Use resources instead
	resource?: string;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [AzureLogsQueryBuilder](./builder-AzureLogsQueryBuilder.md)
