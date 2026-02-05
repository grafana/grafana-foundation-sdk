"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.MapViewConfigBuilder = void 0;
const tslib_1 = require("tslib");
const geomap = tslib_1.__importStar(require("../geomap"));
class MapViewConfigBuilder {
    constructor() {
        this.internal = geomap.defaultMapViewConfig();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    id(id) {
        this.internal.id = id;
        return this;
    }
    lat(lat) {
        this.internal.lat = lat;
        return this;
    }
    lon(lon) {
        this.internal.lon = lon;
        return this;
    }
    zoom(zoom) {
        this.internal.zoom = zoom;
        return this;
    }
    minZoom(minZoom) {
        this.internal.minZoom = minZoom;
        return this;
    }
    maxZoom(maxZoom) {
        this.internal.maxZoom = maxZoom;
        return this;
    }
    padding(padding) {
        this.internal.padding = padding;
        return this;
    }
    allLayers(allLayers) {
        this.internal.allLayers = allLayers;
        return this;
    }
    lastOnly(lastOnly) {
        this.internal.lastOnly = lastOnly;
        return this;
    }
    layer(layer) {
        this.internal.layer = layer;
        return this;
    }
    shared(shared) {
        this.internal.shared = shared;
        return this;
    }
}
exports.MapViewConfigBuilder = MapViewConfigBuilder;
//# sourceMappingURL=mapViewConfigBuilder.gen.js.map