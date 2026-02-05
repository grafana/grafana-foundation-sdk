import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';
export declare class TermsBuilder implements cog.Builder<elasticsearch.Terms> {
    protected readonly internal: elasticsearch.Terms;
    constructor();
    /**
     * Builds the object.
     */
    build(): elasticsearch.Terms;
    field(field: string): this;
    id(id: string): this;
    settings(settings: {
        order?: elasticsearch.TermsOrder;
        size?: string;
        min_doc_count?: string;
        orderBy?: string;
        missing?: string;
    }): this;
}
