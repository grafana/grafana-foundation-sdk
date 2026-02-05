"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.ConditionalRenderingDataSpecBuilder = void 0;
const tslib_1 = require("tslib");
const dashboardv2beta1 = tslib_1.__importStar(require("../dashboardv2beta1"));
class ConditionalRenderingDataSpecBuilder {
    constructor() {
        this.internal = dashboardv2beta1.defaultConditionalRenderingDataSpec();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    value(value) {
        this.internal.value = value;
        return this;
    }
}
exports.ConditionalRenderingDataSpecBuilder = ConditionalRenderingDataSpecBuilder;
//# sourceMappingURL=conditionalRenderingDataSpecBuilder.gen.js.map