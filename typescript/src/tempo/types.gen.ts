// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as dashboard from '../dashboard';


export interface TempoQuery {
	// A unique identifier for the query within the list of targets.
	// In server side expressions, the refId is used as a variable name to identify results.
	// By default, the UI will assign A->Z; however setting meaningful names may be useful.
	refId: string;
	// If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
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
	// Filters to be included in a PromQL query to select data for the service graph. Example: {client="app",service="app"}. Providing multiple values will produce union of results for each filter, using PromQL OR operator internally.
	serviceMapQuery?: string | string[];
	// Use service.namespace in addition to service.name to uniquely identify a service.
	serviceMapIncludeNamespace?: boolean;
	// Defines the maximum number of traces that are returned from Tempo
	limit?: number;
	// Defines the maximum number of spans per spanset that are returned from Tempo
	spss?: number;
	filters: TraceqlFilter[];
	// Filters that are used to query the metrics summary
	groupBy?: TraceqlFilter[];
	// The type of the table that is used to display the search results
	tableType?: SearchTableType;
	// For mixed data sources the selected datasource is on the query level.
	// For non mixed scenarios this is undefined.
	// TODO find a better way to do this ^ that's friendly to schema
	// TODO this shouldn't be unknown but DataSourceRef | null
	datasource?: dashboard.DataSourceRef;
	// For metric queries, the step size to use
	step?: string;
	_implementsDataqueryVariant(): void;
}

export const defaultTempoQuery = (): TempoQuery => ({
	refId: "",
	filters: [],
	_implementsDataqueryVariant: () => {},
});

export enum TempoQueryType {
	Traceql = "traceql",
	TraceqlSearch = "traceqlSearch",
	ServiceMap = "serviceMap",
	Upload = "upload",
	NativeSearch = "nativeSearch",
	TraceId = "traceId",
	Clear = "clear",
}

export const defaultTempoQueryType = (): TempoQueryType => (TempoQueryType.Traceql);

// The state of the TraceQL streaming search query
export enum SearchStreamingState {
	Pending = "pending",
	Streaming = "streaming",
	Done = "done",
	Error = "error",
}

export const defaultSearchStreamingState = (): SearchStreamingState => (SearchStreamingState.Pending);

// The type of the table that is used to display the search results
export enum SearchTableType {
	Traces = "traces",
	Spans = "spans",
	Raw = "raw",
}

export const defaultSearchTableType = (): SearchTableType => (SearchTableType.Traces);

// static fields are pre-set in the UI, dynamic fields are added by the user
export enum TraceqlSearchScope {
	Intrinsic = "intrinsic",
	Unscoped = "unscoped",
	Event = "event",
	Instrumentation = "instrumentation",
	Link = "link",
	Resource = "resource",
	Span = "span",
}

export const defaultTraceqlSearchScope = (): TraceqlSearchScope => (TraceqlSearchScope.Intrinsic);

export interface TraceqlFilter {
	// Uniquely identify the filter, will not be used in the query generation
	id: string;
	// The tag for the search filter, for example: .http.status_code, .service.name, status
	tag?: string;
	// The operator that connects the tag to the value, for example: =, >, !=, =~
	operator?: string;
	// The value for the search filter
	value?: string | string[];
	// The type of the value, used for example to check whether we need to wrap the value in quotes when generating the query
	valueType?: string;
	// The scope of the filter, can either be unscoped/all scopes, resource or span
	scope?: TraceqlSearchScope;
}

export const defaultTraceqlFilter = (): TraceqlFilter => ({
	id: "",
});

