"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.FrameGeometrySourceBuilder = void 0;
const tslib_1 = require("tslib");
const common = tslib_1.__importStar(require("../common"));
class FrameGeometrySourceBuilder {
    constructor() {
        this.internal = common.defaultFrameGeometrySource();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    mode(mode) {
        this.internal.mode = mode;
        return this;
    }
    // Field mappings
    geohash(geohash) {
        this.internal.geohash = geohash;
        return this;
    }
    latitude(latitude) {
        this.internal.latitude = latitude;
        return this;
    }
    longitude(longitude) {
        this.internal.longitude = longitude;
        return this;
    }
    wkt(wkt) {
        this.internal.wkt = wkt;
        return this;
    }
    lookup(lookup) {
        this.internal.lookup = lookup;
        return this;
    }
    // Path to Gazetteer
    gazetteer(gazetteer) {
        this.internal.gazetteer = gazetteer;
        return this;
    }
}
exports.FrameGeometrySourceBuilder = FrameGeometrySourceBuilder;
//# sourceMappingURL=frameGeometrySourceBuilder.gen.js.map