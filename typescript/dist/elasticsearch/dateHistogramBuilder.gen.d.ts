import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';
export declare class DateHistogramBuilder implements cog.Builder<elasticsearch.DateHistogram> {
    protected readonly internal: elasticsearch.DateHistogram;
    constructor();
    /**
     * Builds the object.
     */
    build(): elasticsearch.DateHistogram;
    field(field: string): this;
    id(id: string): this;
    settings(settings: {
        interval?: string;
        min_doc_count?: string;
        trimEdges?: string;
        offset?: string;
        timeZone?: string;
    }): this;
}
