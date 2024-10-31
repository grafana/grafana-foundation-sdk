---
title: <span class="badge object-type-interface"></span> CloudMonitoringQuery
---
# <span class="badge object-type-interface"></span> CloudMonitoringQuery

## Definition

```typescript
export interface CloudMonitoringQuery {
	// A unique identifier for the query within the list of targets.
	// In server side expressions, the refId is used as a variable name to identify results.
	// By default, the UI will assign A->Z; however setting meaningful names may be useful.
	refId: string;
	// true if query is disabled (ie should not be returned to the dashboard)
	// Note this does not always imply that the query should not be executed since
	// the results from a hidden query may be used as the input to other queries (SSE etc)
	hide?: boolean;
	// Specify the query flavor
	// TODO make this required and give it a default
	queryType?: string;
	// Aliases can be set to modify the legend labels. e.g. {{metric.label.xxx}}. See docs for more detail.
	aliasBy?: string;
	// GCM query type.
	// queryType: #QueryType
	// Time Series List sub-query properties.
	timeSeriesList?: googlecloudmonitoring.TimeSeriesList;
	// Time Series sub-query properties.
	timeSeriesQuery?: googlecloudmonitoring.TimeSeriesQuery;
	// SLO sub-query properties.
	sloQuery?: googlecloudmonitoring.SLOQuery;
	// For mixed data sources the selected datasource is on the query level.
	// For non mixed scenarios this is undefined.
	// TODO find a better way to do this ^ that's friendly to schema
	// TODO this shouldn't be unknown but DataSourceRef | null
	datasource?: dashboard.DataSourceRef;
	// Time interval in milliseconds.
	intervalMs?: number;
	_implementsDataqueryVariant(): void;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [CloudMonitoringQueryBuilder](./builder-CloudMonitoringQueryBuilder.md)
