"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.HeatmapCalculationOptionsBuilder = void 0;
const tslib_1 = require("tslib");
const common = tslib_1.__importStar(require("../common"));
class HeatmapCalculationOptionsBuilder {
    constructor() {
        this.internal = common.defaultHeatmapCalculationOptions();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    // The number of buckets to use for the xAxis in the heatmap
    xBuckets(xBuckets) {
        const xBucketsResource = xBuckets.build();
        this.internal.xBuckets = xBucketsResource;
        return this;
    }
    // The number of buckets to use for the yAxis in the heatmap
    yBuckets(yBuckets) {
        const yBucketsResource = yBuckets.build();
        this.internal.yBuckets = yBucketsResource;
        return this;
    }
}
exports.HeatmapCalculationOptionsBuilder = HeatmapCalculationOptionsBuilder;
//# sourceMappingURL=heatmapCalculationOptionsBuilder.gen.js.map