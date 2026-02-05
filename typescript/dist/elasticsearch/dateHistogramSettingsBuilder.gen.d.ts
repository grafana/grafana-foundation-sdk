import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';
export declare class DateHistogramSettingsBuilder implements cog.Builder<elasticsearch.DateHistogramSettings> {
    protected readonly internal: elasticsearch.DateHistogramSettings;
    constructor();
    /**
     * Builds the object.
     */
    build(): elasticsearch.DateHistogramSettings;
    interval(interval: string): this;
    minDocCount(minDocCount: string): this;
    trimEdges(trimEdges: string): this;
    offset(offset: string): this;
    timeZone(timeZone: string): this;
}
