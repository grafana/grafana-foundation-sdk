import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';
export declare class MetricAggregationWithMissingSupportBuilder implements cog.Builder<elasticsearch.MetricAggregationWithMissingSupport> {
    protected readonly internal: elasticsearch.MetricAggregationWithMissingSupport;
    constructor();
    /**
     * Builds the object.
     */
    build(): elasticsearch.MetricAggregationWithMissingSupport;
    settings(settings: {
        missing?: string;
    }): this;
    type(type: elasticsearch.MetricAggregationType): this;
    id(id: string): this;
    hide(hide: boolean): this;
}
