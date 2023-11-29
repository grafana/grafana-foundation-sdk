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
	expr?: string;
	legendFormat?: string;
	maxLines?: number;
	resolution?: number;
	editorMode?: QueryEditorMode;
	range?: boolean;
	instant?: boolean;
	step?: string;
	refId?: string;
	hide?: boolean;
	queryType?: string;
	datasource?: any;
	_implementsDataqueryVariant(): void;
}

export const defaultDataquery = (): dataquery => ({
	_implementsDataqueryVariant: () => {},
});

