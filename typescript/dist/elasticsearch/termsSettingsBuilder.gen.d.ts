import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';
export declare class TermsSettingsBuilder implements cog.Builder<elasticsearch.TermsSettings> {
    protected readonly internal: elasticsearch.TermsSettings;
    constructor();
    /**
     * Builds the object.
     */
    build(): elasticsearch.TermsSettings;
    order(order: elasticsearch.TermsOrder): this;
    size(size: string): this;
    minDocCount(minDocCount: string): this;
    orderBy(orderBy: string): this;
    missing(missing: string): this;
}
