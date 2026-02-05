"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.HeatmapCalculationBucketConfigBuilder = void 0;
const tslib_1 = require("tslib");
const common = tslib_1.__importStar(require("../common"));
class HeatmapCalculationBucketConfigBuilder {
    constructor() {
        this.internal = common.defaultHeatmapCalculationBucketConfig();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    // Sets the bucket calculation mode
    mode(mode) {
        this.internal.mode = mode;
        return this;
    }
    // The number of buckets to use for the axis in the heatmap
    value(value) {
        this.internal.value = value;
        return this;
    }
    // Controls the scale of the buckets
    scale(scale) {
        const scaleResource = scale.build();
        this.internal.scale = scaleResource;
        return this;
    }
}
exports.HeatmapCalculationBucketConfigBuilder = HeatmapCalculationBucketConfigBuilder;
//# sourceMappingURL=heatmapCalculationBucketConfigBuilder.gen.js.map