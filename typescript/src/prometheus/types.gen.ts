// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as common from '../common';


export interface AdhocFilters {
	key: string;
	operator: string;
	value: string;
	values?: string[];
}

export const defaultAdhocFilters = (): AdhocFilters => ({
	key: "",
	operator: "",
	value: "",
});

export interface ResultAssertions {
	// Maximum frame count
	maxFrames?: number;
	// Type asserts that the frame matches a known type structure.
	// Possible enum values:
	//  - `""` 
	//  - `"timeseries-wide"` 
	//  - `"timeseries-long"` 
	//  - `"timeseries-many"` 
	//  - `"timeseries-multi"` 
	//  - `"directory-listing"` 
	//  - `"table"` 
	//  - `"numeric-wide"` 
	//  - `"numeric-multi"` 
	//  - `"numeric-long"` 
	//  - `"log-lines"` 
	type?: ResultAssertionsType;
	// TypeVersion is the version of the Type property. Versions greater than 0.0 correspond to the dataplane
	// contract documentation https://grafana.github.io/dataplane/contract/.
	typeVersion: number[];
}

export const defaultResultAssertions = (): ResultAssertions => ({
	typeVersion: [],
});

export interface ScopesFilters {
	key: string;
	operator: string;
	value: string;
	values?: string[];
}

export const defaultScopesFilters = (): ScopesFilters => ({
	key: "",
	operator: "",
	value: "",
});

export interface Scopes {
	defaultPath?: string[];
	filters?: ScopesFilters[];
	title: string;
}

export const defaultScopes = (): Scopes => ({
	title: "",
});

export interface TimeRange {
	// From is the start time of the query.
	from: string;
	// To is the end time of the query.
	to: string;
}

export const defaultTimeRange = (): TimeRange => ({
	from: "now-6h",
	to: "now",
});

export interface Dataquery {
	// Additional Ad-hoc filters that take precedence over Scope on conflict.
	adhocFilters?: AdhocFilters[];
	// The datasource
	datasource?: common.DataSourceRef;
	// what we should show in the editor
	// Possible enum values:
	//  - `"builder"` 
	//  - `"code"` 
	editorMode?: QueryEditorMode;
	// Execute an additional query to identify interesting raw samples relevant for the given expr
	exemplar?: boolean;
	// The actual expression/query that will be evaluated by Prometheus
	expr: string;
	// The response format
	// Possible enum values:
	//  - `"time_series"` 
	//  - `"table"` 
	//  - `"heatmap"` 
	format?: PromQueryFormat;
	// Group By parameters to apply to aggregate expressions in the query
	groupByKeys?: string[];
	// true if query is disabled (ie should not be returned to the dashboard)
	// NOTE: this does not always imply that the query should not be executed since
	// the results from a hidden query may be used as the input to other queries (SSE etc)
	hide?: boolean;
	// Returns only the latest value that Prometheus has scraped for the requested time series
	instant?: boolean;
	// Used to specify how many times to divide max data points by. We use max data points under query options
	// See https://github.com/grafana/grafana/issues/48081
	// Deprecated: use interval
	intervalFactor?: number;
	// Interval is the suggested duration between time points in a time series query.
	// NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated
	// from the interval required to fill a pixels in the visualization
	intervalMs?: number;
	// Series name override or template. Ex. {{hostname}} will be replaced with label value for hostname
	legendFormat?: string;
	// MaxDataPoints is the maximum number of data points that should be returned from a time series query.
	// NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated
	// from the number of pixels visible in a visualization
	maxDataPoints?: number;
	// QueryType is an optional identifier for the type of query.
	// It can be used to distinguish different types of queries.
	queryType?: string;
	// Returns a Range vector, comprised of a set of time series containing a range of data points over time for each time series
	range?: boolean;
	// RefID is the unique identifier of the query, set by the frontend call.
	refId?: string;
	// Optionally define expected query result behavior
	resultAssertions?: ResultAssertions;
	// A set of filters applied to apply to the query
	scopes?: Scopes[];
	// TimeRange represents the query range
	// NOTE: unlike generic /ds/query, we can now send explicit time values in each query
	// NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
	timeRange?: TimeRange;
	// An additional lower limit for the step parameter of the Prometheus query and for the
	// `$__interval` and `$__rate_interval` variables.
	interval?: string;
	_implementsDataqueryVariant(): void;
}

export const defaultDataquery = (): Dataquery => ({
	expr: "",
	_implementsDataqueryVariant: () => {},
});

export enum ResultAssertionsType {
	None = "",
	TimeseriesWide = "timeseries-wide",
	TimeseriesLong = "timeseries-long",
	TimeseriesMany = "timeseries-many",
	TimeseriesMulti = "timeseries-multi",
	DirectoryListing = "directory-listing",
	Table = "table",
	NumericWide = "numeric-wide",
	NumericMulti = "numeric-multi",
	NumericLong = "numeric-long",
	LogLines = "log-lines",
}

export const defaultResultAssertionsType = (): ResultAssertionsType => (ResultAssertionsType.None);

export enum QueryEditorMode {
	Builder = "builder",
	Code = "code",
}

export const defaultQueryEditorMode = (): QueryEditorMode => (QueryEditorMode.Builder);

export enum PromQueryFormat {
	TimeSeries = "time_series",
	Table = "table",
	Heatmap = "heatmap",
}

export const defaultPromQueryFormat = (): PromQueryFormat => (PromQueryFormat.TimeSeries);

