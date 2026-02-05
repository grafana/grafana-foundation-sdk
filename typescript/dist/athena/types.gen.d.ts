import * as common from '../common';
export interface Dataquery {
    format: FormatOptions;
    connectionArgs: ConnectionArgs;
    table?: string;
    column?: string;
    queryID?: string;
    refId?: string;
    hide?: boolean;
    queryType?: string;
    rawSQL: string;
    datasource?: common.DataSourceRef;
    _implementsDataqueryVariant(): void;
}
export declare const defaultDataquery: () => Dataquery;
export declare enum FormatOptions {
    TimeSeries = 0,
    Table = 1,
    Logs = 2
}
export declare const defaultFormatOptions: () => FormatOptions;
export interface ConnectionArgs {
    region?: string;
    catalog?: string;
    database?: string;
    resultReuseEnabled?: boolean;
    resultReuseMaxAgeInMinutes?: number;
}
export declare const defaultConnectionArgs: () => ConnectionArgs;
export declare const defaultKey = "__default";
