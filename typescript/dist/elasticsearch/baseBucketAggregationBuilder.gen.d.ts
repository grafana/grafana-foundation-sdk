import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';
export declare class BaseBucketAggregationBuilder implements cog.Builder<elasticsearch.BaseBucketAggregation> {
    protected readonly internal: elasticsearch.BaseBucketAggregation;
    constructor();
    /**
     * Builds the object.
     */
    build(): elasticsearch.BaseBucketAggregation;
    id(id: string): this;
    type(type: elasticsearch.BucketAggregationType): this;
    settings(settings: any): this;
}
