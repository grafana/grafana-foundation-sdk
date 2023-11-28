// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';

export class GeoHashGridSettingsBuilder implements cog.OptionsBuilder<elasticsearch.GeoHashGridSettings> {
    private readonly internal: elasticsearch.GeoHashGridSettings;

    constructor() {
        this.internal = elasticsearch.defaultGeoHashGridSettings();
    }

    build(): elasticsearch.GeoHashGridSettings {
        return this.internal;
    }

    precision(precision: string): this {
        this.internal.precision = precision;
        return this;
    }
}
