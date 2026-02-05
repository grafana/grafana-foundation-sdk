"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.HeatmapColorOptionsBuilder = void 0;
const tslib_1 = require("tslib");
const heatmap = tslib_1.__importStar(require("../heatmap"));
// Controls various color options
class HeatmapColorOptionsBuilder {
    constructor() {
        this.internal = heatmap.defaultHeatmapColorOptions();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    // Sets the color mode
    mode(mode) {
        this.internal.mode = mode;
        return this;
    }
    // Controls the color scheme used
    scheme(scheme) {
        this.internal.scheme = scheme;
        return this;
    }
    // Controls the color fill when in opacity mode
    fill(fill) {
        this.internal.fill = fill;
        return this;
    }
    // Controls the color scale
    scale(scale) {
        this.internal.scale = scale;
        return this;
    }
    // Controls the exponent when scale is set to exponential
    exponent(exponent) {
        this.internal.exponent = exponent;
        return this;
    }
    // Controls the number of color steps
    steps(steps) {
        if (!(steps >= 2)) {
            throw new Error("steps must be >= 2");
        }
        if (!(steps <= 128)) {
            throw new Error("steps must be <= 128");
        }
        this.internal.steps = steps;
        return this;
    }
    // Reverses the color scheme
    reverse(reverse) {
        this.internal.reverse = reverse;
        return this;
    }
    // Sets the minimum value for the color scale
    min(min) {
        this.internal.min = min;
        return this;
    }
    // Sets the maximum value for the color scale
    max(max) {
        this.internal.max = max;
        return this;
    }
}
exports.HeatmapColorOptionsBuilder = HeatmapColorOptionsBuilder;
//# sourceMappingURL=heatmapColorOptionsBuilder.gen.js.map