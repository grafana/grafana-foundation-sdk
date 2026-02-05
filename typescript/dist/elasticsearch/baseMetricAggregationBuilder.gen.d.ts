import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';
export declare class BaseMetricAggregationBuilder implements cog.Builder<elasticsearch.BaseMetricAggregation> {
    protected readonly internal: elasticsearch.BaseMetricAggregation;
    constructor();
    /**
     * Builds the object.
     */
    build(): elasticsearch.BaseMetricAggregation;
    type(type: elasticsearch.MetricAggregationType): this;
    id(id: string): this;
    hide(hide: boolean): this;
}
