import * as common from '../common';
export declare enum PyroscopeQueryType {
    Metrics = "metrics",
    Profile = "profile",
    Both = "both"
}
export declare const defaultPyroscopeQueryType: () => PyroscopeQueryType;
export interface dataquery {
    labelSelector: string;
    spanSelector?: string[];
    profileTypeId: string;
    groupBy: string[];
    limit?: number;
    maxNodes?: number;
    refId?: string;
    hide?: boolean;
    queryType?: string;
    datasource?: common.DataSourceRef;
    _implementsDataqueryVariant(): void;
}
export declare const defaultDataquery: () => dataquery;
