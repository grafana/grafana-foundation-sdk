// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';

export class GeoHashGridSettingsBuilder implements cog.Builder<elasticsearch.GeoHashGridSettings> {
    protected readonly internal: elasticsearch.GeoHashGridSettings;

    constructor() {
        this.internal = elasticsearch.defaultGeoHashGridSettings();
    }

    /**
     * Builds the object.
     */
    build(): elasticsearch.GeoHashGridSettings {
        return this.internal;
    }

    precision(precision: string): this {
        this.internal.precision = precision;
        return this;
    }
}
