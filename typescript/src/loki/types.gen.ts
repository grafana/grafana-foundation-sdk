// Code generated - EDITING IS FUTILE. DO NOT EDIT.

export enum QueryEditorMode {
	Code = "code",
	Builder = "builder",
}

export const defaultQueryEditorMode = (): QueryEditorMode => (QueryEditorMode.Code);

export enum LokiQueryType {
	Range = "range",
	Instant = "instant",
	Stream = "stream",
}

export const defaultLokiQueryType = (): LokiQueryType => (LokiQueryType.Range);

export enum SupportingQueryType {
	LogsVolume = "logsVolume",
	LogsSample = "logsSample",
	DataSample = "dataSample",
}

export const defaultSupportingQueryType = (): SupportingQueryType => (SupportingQueryType.LogsVolume);

export enum LokiQueryDirection {
	Forward = "forward",
	Backward = "backward",
}

export const defaultLokiQueryDirection = (): LokiQueryDirection => (LokiQueryDirection.Forward);

export interface dataquery {
	// The LogQL query.
	expr: string;
	// Used to override the name of the series.
	legendFormat?: string;
	// Used to limit the number of log rows returned.
	maxLines?: number;
	// @deprecated, now use step.
	resolution?: number;
	editorMode?: QueryEditorMode;
	// @deprecated, now use queryType.
	range?: boolean;
	// @deprecated, now use queryType.
	instant?: boolean;
	// Used to set step value for range queries.
	step?: string;
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
	datasource?: any;
	_implementsDataqueryVariant(): void;
}

export const defaultDataquery = (): dataquery => ({
	expr: "",
	refId: "",
	_implementsDataqueryVariant: () => {},
});

