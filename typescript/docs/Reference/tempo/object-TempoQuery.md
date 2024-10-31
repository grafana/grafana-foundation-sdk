---
title: <span class="badge object-type-interface"></span> TempoQuery
---
# <span class="badge object-type-interface"></span> TempoQuery

## Definition

```typescript
export interface TempoQuery {
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
	// TraceQL query or trace ID
	query?: string;
	// @deprecated Logfmt query to filter traces by their tags. Example: http.status_code=200 error=true
	search?: string;
	// @deprecated Query traces by service name
	serviceName?: string;
	// @deprecated Query traces by span name
	spanName?: string;
	// @deprecated Define the minimum duration to select traces. Use duration format, for example: 1.2s, 100ms
	minDuration?: string;
	// @deprecated Define the maximum duration to select traces. Use duration format, for example: 1.2s, 100ms
	maxDuration?: string;
	// Filters to be included in a PromQL query to select data for the service graph. Example: {client="app",service="app"}
	serviceMapQuery?: string;
	// Use service.namespace in addition to service.name to uniquely identify a service.
	serviceMapIncludeNamespace?: boolean;
	// Defines the maximum number of traces that are returned from Tempo
	limit?: number;
	// Defines the maximum number of spans per spanset that are returned from Tempo
	spss?: number;
	filters: tempo.TraceqlFilter[];
	// Filters that are used to query the metrics summary
	groupBy?: tempo.TraceqlFilter[];
	// For mixed data sources the selected datasource is on the query level.
	// For non mixed scenarios this is undefined.
	// TODO find a better way to do this ^ that's friendly to schema
	// TODO this shouldn't be unknown but DataSourceRef | null
	datasource?: dashboard.DataSourceRef;
	// The type of the table that is used to display the search results
	tableType?: tempo.SearchTableType;
	_implementsDataqueryVariant(): void;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [TempoQueryBuilder](./builder-TempoQueryBuilder.md)
