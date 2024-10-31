---
title: <span class="badge object-type-interface"></span> AzureTracesQuery
---
# <span class="badge object-type-interface"></span> AzureTracesQuery

Application Insights Traces sub-query properties

## Definition

```typescript
export interface AzureTracesQuery {
	// Specifies the format results should be returned as.
	resultFormat?: azuremonitor.ResultFormat;
	// Array of resource URIs to be queried.
	resources?: string[];
	// Operation ID. Used only for Traces queries.
	operationId?: string;
	// Types of events to filter by.
	traceTypes?: string[];
	// Filters for property values.
	filters?: azuremonitor.AzureTracesFilter[];
	// KQL query to be executed.
	query?: string;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [AzureTracesQueryBuilder](./builder-AzureTracesQueryBuilder.md)
