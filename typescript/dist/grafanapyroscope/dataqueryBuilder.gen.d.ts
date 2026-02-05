import * as cog from '../cog';
import * as grafanapyroscope from '../grafanapyroscope';
import * as common from '../common';
export declare class DataqueryBuilder implements cog.Builder<cog.Dataquery> {
    protected readonly internal: grafanapyroscope.dataquery;
    constructor();
    /**
     * Builds the object.
     */
    build(): grafanapyroscope.dataquery;
    labelSelector(labelSelector: string): this;
    spanSelector(spanSelector: string[]): this;
    profileTypeId(profileTypeId: string): this;
    groupBy(groupBy: string[]): this;
    limit(limit: number): this;
    maxNodes(maxNodes: number): this;
    refId(refId: string): this;
    hide(hide: boolean): this;
    queryType(queryType: string): this;
    datasource(datasource: common.DataSourceRef): this;
}
