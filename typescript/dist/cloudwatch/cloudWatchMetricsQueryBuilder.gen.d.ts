import * as cog from '../cog';
import * as cloudwatch from '../cloudwatch';
import * as common from '../common';
export declare class CloudWatchMetricsQueryBuilder implements cog.Builder<cog.Dataquery> {
    protected readonly internal: cloudwatch.CloudWatchMetricsQuery;
    constructor();
    /**
     * Builds the object.
     */
    build(): cloudwatch.CloudWatchMetricsQuery;
    queryMode(queryMode: cloudwatch.CloudWatchQueryMode): this;
    metricQueryType(metricQueryType: cloudwatch.MetricQueryType): this;
    metricEditorMode(metricEditorMode: cloudwatch.MetricEditorMode): this;
    id(id: string): this;
    alias(alias: string): this;
    label(label: string): this;
    expression(expression: string): this;
    sqlExpression(sqlExpression: string): this;
    refId(refId: string): this;
    hide(hide: boolean): this;
    queryType(queryType: string): this;
    region(region: string): this;
    namespace(namespace: string): this;
    metricName(metricName: string): this;
    dimensions(dimensions: cloudwatch.Dimensions): this;
    matchExact(matchExact: boolean): this;
    period(period: string): this;
    accountId(accountId: string): this;
    statistic(statistic: string): this;
    sql(sql: cog.Builder<cloudwatch.SQLExpression>): this;
    datasource(datasource: common.DataSourceRef): this;
    statistics(statistics: string[]): this;
}
