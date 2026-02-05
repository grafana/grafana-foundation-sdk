import * as common from '../common';
export declare enum QueryEditorMode {
    Code = "code",
    Builder = "builder"
}
export declare const defaultQueryEditorMode: () => QueryEditorMode;
export declare enum PromQueryFormat {
    TimeSeries = "time_series",
    Table = "table",
    Heatmap = "heatmap"
}
export declare const defaultPromQueryFormat: () => PromQueryFormat;
export interface dataquery {
    expr: string;
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
    datasource?: common.DataSourceRef;
    interval?: string;
    _implementsDataqueryVariant(): void;
}
export declare const defaultDataquery: () => dataquery;
