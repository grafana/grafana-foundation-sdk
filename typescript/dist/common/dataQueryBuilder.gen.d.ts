import * as cog from '../cog';
import * as common from '../common';
export declare class DataQueryBuilder implements cog.Builder<common.DataQuery> {
    protected readonly internal: common.DataQuery;
    constructor();
    /**
     * Builds the object.
     */
    build(): common.DataQuery;
    refId(refId: string): this;
    hide(hide: boolean): this;
    queryType(queryType: string): this;
    datasource(datasource: any): this;
}
