// Code generated - EDITING IS FUTILE. DO NOT EDIT.

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
	expr?: string;
	instant?: boolean;
	range?: boolean;
	exemplar?: boolean;
	editorMode?: QueryEditorMode;
	format?: PromQueryFormat;
	legendFormat?: string;
	intervalFactor?: number;
	scope?: {
		matchers: string;
	};
	refId?: string;
	hide?: boolean;
	queryType?: string;
	datasource?: any;
	// An additional lower limit for the step parameter of the Prometheus query and for the
	// `$__interval` and `$__rate_interval` variables.
	interval?: string;
	_implementsDataqueryVariant(): void;
}

export const defaultDataquery = (): dataquery => ({
	_implementsDataqueryVariant: () => {},
});

