import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';
export declare class GeoHashGridSettingsBuilder implements cog.Builder<elasticsearch.GeoHashGridSettings> {
    protected readonly internal: elasticsearch.GeoHashGridSettings;
    constructor();
    /**
     * Builds the object.
     */
    build(): elasticsearch.GeoHashGridSettings;
    precision(precision: string): this;
}
