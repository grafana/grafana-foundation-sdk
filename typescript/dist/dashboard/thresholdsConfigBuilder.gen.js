"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.ThresholdsConfigBuilder = void 0;
const tslib_1 = require("tslib");
const dashboard = tslib_1.__importStar(require("../dashboard"));
// Thresholds configuration for the panel
class ThresholdsConfigBuilder {
    constructor() {
        this.internal = dashboard.defaultThresholdsConfig();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    // Thresholds mode.
    mode(mode) {
        this.internal.mode = mode;
        return this;
    }
    // Must be sorted by 'value', first value is always -Infinity
    steps(steps) {
        this.internal.steps = steps;
        return this;
    }
}
exports.ThresholdsConfigBuilder = ThresholdsConfigBuilder;
//# sourceMappingURL=thresholdsConfigBuilder.gen.js.map