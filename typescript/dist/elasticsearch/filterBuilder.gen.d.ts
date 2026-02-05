import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';
export declare class FilterBuilder implements cog.Builder<elasticsearch.Filter> {
    protected readonly internal: elasticsearch.Filter;
    constructor();
    /**
     * Builds the object.
     */
    build(): elasticsearch.Filter;
    query(query: string): this;
    label(label: string): this;
}
