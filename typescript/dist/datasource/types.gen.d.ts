import * as common from '../common';
export interface Dataquery {
    panelId: number;
    refId?: string;
    hide?: boolean;
    queryType?: string;
    withTransforms: boolean;
    datasource?: common.DataSourceRef;
    _implementsDataqueryVariant(): void;
}
export declare const defaultDataquery: () => Dataquery;
