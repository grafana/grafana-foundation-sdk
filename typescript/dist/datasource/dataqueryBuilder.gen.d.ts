import * as cog from '../cog';
import * as datasource from '../datasource';
import * as common from '../common';
export declare class DataqueryBuilder implements cog.Builder<cog.Dataquery> {
    protected readonly internal: datasource.Dataquery;
    constructor();
    /**
     * Builds the object.
     */
    build(): datasource.Dataquery;
    panelId(panelId: number): this;
    refId(refId: string): this;
    hide(hide: boolean): this;
    withTransforms(withTransforms: boolean): this;
    datasource(datasource: common.DataSourceRef): this;
}
