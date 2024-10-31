---
title: <span class="badge object-type-interface"></span> TimeSeriesQuery
---
# <span class="badge object-type-interface"></span> TimeSeriesQuery

Time Series sub-query properties.

## Definition

```typescript
export interface TimeSeriesQuery {
	// GCP project to execute the query against.
	projectName: string;
	// MQL query to be executed.
	query: string;
	// To disable the graphPeriod, it should explictly be set to 'disabled'.
	graphPeriod?: string;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [TimeSeriesQueryBuilder](./builder-TimeSeriesQueryBuilder.md)
