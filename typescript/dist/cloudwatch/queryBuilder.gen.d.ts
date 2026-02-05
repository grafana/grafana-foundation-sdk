import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
import * as cloudwatch from '../cloudwatch';
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
    cloudWatchMetricsQuery(cloudWatchMetricsQuery: cog.Builder<cloudwatch.CloudWatchMetricsQuery>): this;
    cloudWatchLogsQuery(cloudWatchLogsQuery: cog.Builder<cloudwatch.CloudWatchLogsQuery>): this;
    cloudWatchAnnotationQuery(cloudWatchAnnotationQuery: cog.Builder<cloudwatch.CloudWatchAnnotationQuery>): this;
}
