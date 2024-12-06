// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as common from '../common';

export class FrameGeometrySourceBuilder implements cog.Builder<common.FrameGeometrySource> {
    protected readonly internal: common.FrameGeometrySource;

    constructor() {
        this.internal = common.defaultFrameGeometrySource();
    }

    /**
     * Builds the object.
     */
    build(): common.FrameGeometrySource {
        return this.internal;
    }

    mode(mode: common.FrameGeometrySourceMode): this {
        this.internal.mode = mode;
        return this;
    }

    // Field mappings
    geohash(geohash: string): this {
        this.internal.geohash = geohash;
        return this;
    }

    latitude(latitude: string): this {
        this.internal.latitude = latitude;
        return this;
    }

    longitude(longitude: string): this {
        this.internal.longitude = longitude;
        return this;
    }

    wkt(wkt: string): this {
        this.internal.wkt = wkt;
        return this;
    }

    lookup(lookup: string): this {
        this.internal.lookup = lookup;
        return this;
    }

    // Path to Gazetteer
    gazetteer(gazetteer: string): this {
        this.internal.gazetteer = gazetteer;
        return this;
    }
}
