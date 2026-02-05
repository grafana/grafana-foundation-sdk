import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';
export declare class FiltersSettingsBuilder implements cog.Builder<elasticsearch.FiltersSettings> {
    protected readonly internal: elasticsearch.FiltersSettings;
    constructor();
    /**
     * Builds the object.
     */
    build(): elasticsearch.FiltersSettings;
    filters(filters: cog.Builder<elasticsearch.Filter>[]): this;
}
