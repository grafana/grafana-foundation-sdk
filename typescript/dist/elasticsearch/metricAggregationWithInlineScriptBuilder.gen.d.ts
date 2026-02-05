import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';
export declare class MetricAggregationWithInlineScriptBuilder implements cog.Builder<elasticsearch.MetricAggregationWithInlineScript> {
    protected readonly internal: elasticsearch.MetricAggregationWithInlineScript;
    constructor();
    /**
     * Builds the object.
     */
    build(): elasticsearch.MetricAggregationWithInlineScript;
    settings(settings: {
        script?: elasticsearch.InlineScript;
    }): this;
    type(type: elasticsearch.MetricAggregationType): this;
    id(id: string): this;
    hide(hide: boolean): this;
}
