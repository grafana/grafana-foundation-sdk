---
title: <span class="badge object-type-interface"></span> DataQuery
---
# <span class="badge object-type-interface"></span> DataQuery

These are the common properties available to all queries in all datasources.

Specific implementations will *extend* this interface, adding the required

properties for the given context.

## Definition

```typescript
export interface DataQuery {
	// A unique identifier for the query within the list of targets.
	// In server side expressions, the refId is used as a variable name to identify results.
	// By default, the UI will assign A->Z; however setting meaningful names may be useful.
	refId: string;
	// If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
	hide?: boolean;
	// Specify the query flavor
	// TODO make this required and give it a default
	queryType?: string;
	// For mixed data sources the selected datasource is on the query level.
	// For non mixed scenarios this is undefined.
	// TODO find a better way to do this ^ that's friendly to schema
	// TODO this shouldn't be unknown but DataSourceRef | null
	datasource?: any;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [DataQueryBuilder](./builder-DataQueryBuilder.md)
