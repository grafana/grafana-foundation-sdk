"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.ConditionalRenderingDataBuilder = void 0;
const tslib_1 = require("tslib");
const dashboardv2beta1 = tslib_1.__importStar(require("../dashboardv2beta1"));
class ConditionalRenderingDataBuilder {
    constructor() {
        this.internal = dashboardv2beta1.defaultConditionalRenderingDataKind();
        this.internal.kind = "ConditionalRenderingData";
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    spec(spec) {
        const specResource = spec.build();
        this.internal.spec = specResource;
        return this;
    }
    value(value) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultConditionalRenderingDataSpec();
        }
        this.internal.spec.value = value;
        return this;
    }
}
exports.ConditionalRenderingDataBuilder = ConditionalRenderingDataBuilder;
//# sourceMappingURL=conditionalRenderingDataBuilder.gen.js.map