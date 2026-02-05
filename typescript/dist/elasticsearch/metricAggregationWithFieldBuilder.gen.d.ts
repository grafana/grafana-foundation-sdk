import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';
export declare class MetricAggregationWithFieldBuilder implements cog.Builder<elasticsearch.MetricAggregationWithField> {
    protected readonly internal: elasticsearch.MetricAggregationWithField;
    constructor();
    /**
     * Builds the object.
     */
    build(): elasticsearch.MetricAggregationWithField;
    field(field: string): this;
    type(type: elasticsearch.MetricAggregationType): this;
    id(id: string): this;
    hide(hide: boolean): this;
}
