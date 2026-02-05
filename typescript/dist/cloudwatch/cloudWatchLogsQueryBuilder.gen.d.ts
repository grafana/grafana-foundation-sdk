import * as cog from '../cog';
import * as cloudwatch from '../cloudwatch';
import * as common from '../common';
export declare class CloudWatchLogsQueryBuilder implements cog.Builder<cog.Dataquery> {
    protected readonly internal: cloudwatch.CloudWatchLogsQuery;
    constructor();
    /**
     * Builds the object.
     */
    build(): cloudwatch.CloudWatchLogsQuery;
    queryMode(queryMode: cloudwatch.CloudWatchQueryMode): this;
    id(id: string): this;
    region(region: string): this;
    expression(expression: string): this;
    statsGroups(statsGroups: string[]): this;
    logGroups(logGroups: cog.Builder<cloudwatch.LogGroup>[]): this;
    logGroupNames(logGroupNames: string[]): this;
    refId(refId: string): this;
    hide(hide: boolean): this;
    queryType(queryType: string): this;
    queryLanguage(queryLanguage: cloudwatch.LogsQueryLanguage): this;
    datasource(datasource: common.DataSourceRef): this;
}
