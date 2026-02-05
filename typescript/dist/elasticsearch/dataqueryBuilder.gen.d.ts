import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';
import * as common from '../common';
export declare class DataqueryBuilder implements cog.Builder<cog.Dataquery> {
    protected readonly internal: elasticsearch.dataquery;
    constructor();
    /**
     * Builds the object.
     */
    build(): elasticsearch.dataquery;
    alias(alias: string): this;
    query(query: string): this;
    timeField(timeField: string): this;
    bucketAggs(bucketAggs: cog.Builder<elasticsearch.BucketAggregation>[]): this;
    metrics(metrics: cog.Builder<elasticsearch.MetricAggregation>[]): this;
    refId(refId: string): this;
    hide(hide: boolean): this;
    queryType(queryType: string): this;
    datasource(datasource: common.DataSourceRef): this;
}
