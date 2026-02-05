import * as common from '../common';
export declare enum QueryEditorMode {
    Code = "code",
    Builder = "builder"
}
export declare const defaultQueryEditorMode: () => QueryEditorMode;
export declare enum LokiQueryType {
    Range = "range",
    Instant = "instant",
    Stream = "stream"
}
export declare const defaultLokiQueryType: () => LokiQueryType;
export declare enum SupportingQueryType {
    LogsVolume = "logsVolume",
    LogsSample = "logsSample",
    DataSample = "dataSample",
    InfiniteScroll = "infiniteScroll"
}
export declare const defaultSupportingQueryType: () => SupportingQueryType;
export declare enum LokiQueryDirection {
    Forward = "forward",
    Backward = "backward",
    Scan = "scan"
}
export declare const defaultLokiQueryDirection: () => LokiQueryDirection;
export interface dataquery {
    expr: string;
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
    datasource?: common.DataSourceRef;
    direction?: LokiQueryDirection;
    _implementsDataqueryVariant(): void;
}
export declare const defaultDataquery: () => dataquery;
