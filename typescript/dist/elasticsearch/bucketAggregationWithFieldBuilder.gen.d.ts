import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';
export declare class BucketAggregationWithFieldBuilder implements cog.Builder<elasticsearch.BucketAggregationWithField> {
    protected readonly internal: elasticsearch.BucketAggregationWithField;
    constructor();
    /**
     * Builds the object.
     */
    build(): elasticsearch.BucketAggregationWithField;
    field(field: string): this;
    id(id: string): this;
    type(type: elasticsearch.BucketAggregationType): this;
    settings(settings: any): this;
}
