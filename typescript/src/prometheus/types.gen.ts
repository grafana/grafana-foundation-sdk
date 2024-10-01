// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as dashboard from '../dashboard';


export enum QueryEditorMode {
	Code = "code",
	Builder = "builder",
}

export const defaultQueryEditorMode = (): QueryEditorMode => (QueryEditorMode.Code);

export enum PromQueryFormat {
	TimeSeries = "time_series",
	Table = "table",
	Heatmap = "heatmap",
}

export const defaultPromQueryFormat = (): PromQueryFormat => (PromQueryFormat.TimeSeries);

export interface dataquery {
	// The actual expression/query that will be evaluated by Prometheus
	expr: string;
	// Returns only the latest value that Prometheus has scraped for the requested time series
	instant?: boolean;
	// Returns a Range vector, comprised of a set of time series containing a range of data points over time for each time series
	range?: boolean;
	// Execute an additional query to identify interesting raw samples relevant for the given expr
	exemplar?: boolean;
	// Specifies which editor is being used to prepare the query. It can be "code" or "builder"
	editorMode?: QueryEditorMode;
	// Query format to determine how to display data points in panel. It can be "time_series", "table", "heatmap"
	format?: PromQueryFormat;
	// Series name override or template. Ex. {{hostname}} will be replaced with label value for hostname
	legendFormat?: string;
	// @deprecated Used to specify how many times to divide max data points by. We use max data points under query options
	// See https://github.com/grafana/grafana/issues/48081
	intervalFactor?: number;
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
	// For mixed data sources the selected datasource is on the query level.
	// For non mixed scenarios this is undefined.
	// TODO find a better way to do this ^ that's friendly to schema
	// TODO this shouldn't be unknown but DataSourceRef | null
	datasource?: dashboard.DataSourceRef;
	// An additional lower limit for the step parameter of the Prometheus query and for the
	// `$__interval` and `$__rate_interval` variables.
	interval?: string;
	_implementsDataqueryVariant(): void;
}

export const defaultDataquery = (): dataquery => ({
	expr: "",
	refId: "",
	_implementsDataqueryVariant: () => {},
});

