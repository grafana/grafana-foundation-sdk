import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';
export declare class HistogramBuilder implements cog.Builder<elasticsearch.Histogram> {
    protected readonly internal: elasticsearch.Histogram;
    constructor();
    /**
     * Builds the object.
     */
    build(): elasticsearch.Histogram;
    field(field: string): this;
    id(id: string): this;
    settings(settings: {
        interval?: string;
        min_doc_count?: string;
    }): this;
}
