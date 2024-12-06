// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as geomap from '../geomap';

export class MapViewConfigBuilder implements cog.Builder<geomap.MapViewConfig> {
    protected readonly internal: geomap.MapViewConfig;

    constructor() {
        this.internal = geomap.defaultMapViewConfig();
    }

    /**
     * Builds the object.
     */
    build(): geomap.MapViewConfig {
        return this.internal;
    }

    id(id: string): this {
        this.internal.id = id;
        return this;
    }

    lat(lat: number): this {
        this.internal.lat = lat;
        return this;
    }

    lon(lon: number): this {
        this.internal.lon = lon;
        return this;
    }

    zoom(zoom: number): this {
        this.internal.zoom = zoom;
        return this;
    }

    minZoom(minZoom: number): this {
        this.internal.minZoom = minZoom;
        return this;
    }

    maxZoom(maxZoom: number): this {
        this.internal.maxZoom = maxZoom;
        return this;
    }

    padding(padding: number): this {
        this.internal.padding = padding;
        return this;
    }

    allLayers(allLayers: boolean): this {
        this.internal.allLayers = allLayers;
        return this;
    }

    lastOnly(lastOnly: boolean): this {
        this.internal.lastOnly = lastOnly;
        return this;
    }

    layer(layer: string): this {
        this.internal.layer = layer;
        return this;
    }

    shared(shared: boolean): this {
        this.internal.shared = shared;
        return this;
    }
}
