import * as cog from '../cog';
import * as parca from '../parca';
import * as common from '../common';
export declare class DataqueryBuilder implements cog.Builder<cog.Dataquery> {
    protected readonly internal: parca.dataquery;
    constructor();
    /**
     * Builds the object.
     */
    build(): parca.dataquery;
    labelSelector(labelSelector: string): this;
    profileTypeId(profileTypeId: string): this;
    refId(refId: string): this;
    hide(hide: boolean): this;
    queryType(queryType: string): this;
    datasource(datasource: common.DataSourceRef): this;
}
