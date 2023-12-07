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
	refId?: string;
	hide?: boolean;
	queryType?: string;
	datasource?: any;
	_implementsDataqueryVariant(): void;
}

export const defaultDataquery = (): dataquery => ({
	_implementsDataqueryVariant: () => {},
});
