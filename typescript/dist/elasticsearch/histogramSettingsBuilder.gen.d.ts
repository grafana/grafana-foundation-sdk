import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';
export declare class HistogramSettingsBuilder implements cog.Builder<elasticsearch.HistogramSettings> {
    protected readonly internal: elasticsearch.HistogramSettings;
    constructor();
    /**
     * Builds the object.
     */
    build(): elasticsearch.HistogramSettings;
    interval(interval: string): this;
    minDocCount(minDocCount: string): this;
}
