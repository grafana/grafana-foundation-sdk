"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.AutoGridLayoutItemSpecBuilder = void 0;
const tslib_1 = require("tslib");
const dashboardv2beta1 = tslib_1.__importStar(require("../dashboardv2beta1"));
class AutoGridLayoutItemSpecBuilder {
    constructor() {
        this.internal = dashboardv2beta1.defaultAutoGridLayoutItemSpec();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    element(element) {
        const elementResource = element.build();
        this.internal.element = elementResource;
        return this;
    }
    repeat(repeat) {
        const repeatResource = repeat.build();
        this.internal.repeat = repeatResource;
        return this;
    }
    conditionalRendering(conditionalRendering) {
        const conditionalRenderingResource = conditionalRendering.build();
        this.internal.conditionalRendering = conditionalRenderingResource;
        return this;
    }
}
exports.AutoGridLayoutItemSpecBuilder = AutoGridLayoutItemSpecBuilder;
//# sourceMappingURL=autoGridLayoutItemSpecBuilder.gen.js.map