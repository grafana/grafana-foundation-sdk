---
title: <span class="badge object-type-interface"></span> dataquery
---
# <span class="badge object-type-interface"></span> dataquery

## Definition

```typescript
export interface dataquery {
	// Alias pattern
	alias?: string;
	// Lucene query
	query?: string;
	// Name of time field
	timeField?: string;
	// List of bucket aggregations
	bucketAggs?: elasticsearch.BucketAggregation[];
	// List of metric aggregations
	metrics?: elasticsearch.MetricAggregation[];
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
	datasource?: dashboard.DataSourceRef;
	_implementsDataqueryVariant(): void;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [DataqueryBuilder](./builder-DataqueryBuilder.md)
