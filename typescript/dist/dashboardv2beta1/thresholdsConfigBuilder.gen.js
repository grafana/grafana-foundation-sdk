"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.ThresholdsConfigBuilder = void 0;
const tslib_1 = require("tslib");
const dashboardv2beta1 = tslib_1.__importStar(require("../dashboardv2beta1"));
class ThresholdsConfigBuilder {
    constructor() {
        this.internal = dashboardv2beta1.defaultThresholdsConfig();
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
    steps(steps) {
        this.internal.steps = steps;
        return this;
    }
}
exports.ThresholdsConfigBuilder = ThresholdsConfigBuilder;
//# sourceMappingURL=thresholdsConfigBuilder.gen.js.map