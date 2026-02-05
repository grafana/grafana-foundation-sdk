import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
import * as googlecloudmonitoring from '../googlecloudmonitoring';
export declare class QueryBuilder implements cog.Builder<dashboardv2beta1.DataQueryKind> {
    protected readonly internal: dashboardv2beta1.DataQueryKind;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.DataQueryKind;
    version(version: string): this;
    datasource(ref: {
        name?: string;
    }): this;
    refId(refId: string): this;
    hide(hide: boolean): this;
    queryType(queryType: string): this;
    aliasBy(aliasBy: string): this;
    timeSeriesList(timeSeriesList: cog.Builder<googlecloudmonitoring.TimeSeriesList>): this;
    timeSeriesQuery(timeSeriesQuery: cog.Builder<googlecloudmonitoring.TimeSeriesQuery>): this;
    sloQuery(sloQuery: cog.Builder<googlecloudmonitoring.SLOQuery>): this;
    promQLQuery(promQLQuery: cog.Builder<googlecloudmonitoring.PromQLQuery>): this;
    intervalMs(intervalMs: number): this;
}
