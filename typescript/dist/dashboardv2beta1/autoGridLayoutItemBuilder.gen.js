"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.AutoGridLayoutItemBuilder = void 0;
const tslib_1 = require("tslib");
const dashboardv2beta1 = tslib_1.__importStar(require("../dashboardv2beta1"));
class AutoGridLayoutItemBuilder {
    constructor(name) {
        this.internal = dashboardv2beta1.defaultAutoGridLayoutItemKind();
        this.internal.kind = "AutoGridLayoutItem";
        this.internal.spec.element.kind = "ElementReference";
        this.internal.spec.element.name = name;
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    element(element) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultAutoGridLayoutItemSpec();
        }
        const elementResource = element.build();
        this.internal.spec.element = elementResource;
        return this;
    }
    repeat(repeat) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultAutoGridLayoutItemSpec();
        }
        const repeatResource = repeat.build();
        this.internal.spec.repeat = repeatResource;
        return this;
    }
    conditionalRendering(conditionalRendering) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultAutoGridLayoutItemSpec();
        }
        const conditionalRenderingResource = conditionalRendering.build();
        this.internal.spec.conditionalRendering = conditionalRenderingResource;
        return this;
    }
    name(name) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultAutoGridLayoutItemSpec();
        }
        if (!this.internal.spec.element) {
            this.internal.spec.element = dashboardv2beta1.defaultElementReference();
        }
        this.internal.spec.element.name = name;
        return this;
    }
}
exports.AutoGridLayoutItemBuilder = AutoGridLayoutItemBuilder;
//# sourceMappingURL=autoGridLayoutItemBuilder.gen.js.map