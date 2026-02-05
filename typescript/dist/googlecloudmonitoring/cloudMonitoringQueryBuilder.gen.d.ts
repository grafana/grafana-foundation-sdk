import * as cog from '../cog';
import * as googlecloudmonitoring from '../googlecloudmonitoring';
import * as common from '../common';
export declare class CloudMonitoringQueryBuilder implements cog.Builder<cog.Dataquery> {
    protected readonly internal: googlecloudmonitoring.CloudMonitoringQuery;
    constructor();
    /**
     * Builds the object.
     */
    build(): googlecloudmonitoring.CloudMonitoringQuery;
    refId(refId: string): this;
    hide(hide: boolean): this;
    queryType(queryType: string): this;
    aliasBy(aliasBy: string): this;
    timeSeriesList(timeSeriesList: cog.Builder<googlecloudmonitoring.TimeSeriesList>): this;
    timeSeriesQuery(timeSeriesQuery: cog.Builder<googlecloudmonitoring.TimeSeriesQuery>): this;
    sloQuery(sloQuery: cog.Builder<googlecloudmonitoring.SLOQuery>): this;
    promQLQuery(promQLQuery: cog.Builder<googlecloudmonitoring.PromQLQuery>): this;
    datasource(datasource: common.DataSourceRef): this;
    intervalMs(intervalMs: number): this;
}
