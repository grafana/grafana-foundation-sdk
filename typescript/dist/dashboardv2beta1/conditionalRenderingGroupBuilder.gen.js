"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.ConditionalRenderingGroupBuilder = void 0;
const tslib_1 = require("tslib");
const dashboardv2beta1 = tslib_1.__importStar(require("../dashboardv2beta1"));
class ConditionalRenderingGroupBuilder {
    constructor() {
        this.internal = dashboardv2beta1.defaultConditionalRenderingGroupKind();
        this.internal.kind = "ConditionalRenderingGroup";
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
    visibility(visibility) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultConditionalRenderingGroupSpec();
        }
        this.internal.spec.visibility = visibility;
        return this;
    }
    condition(condition) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultConditionalRenderingGroupSpec();
        }
        this.internal.spec.condition = condition;
        return this;
    }
    items(items) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultConditionalRenderingGroupSpec();
        }
        const itemsResources = items.map(builder1 => builder1.build());
        this.internal.spec.items = itemsResources;
        return this;
    }
}
exports.ConditionalRenderingGroupBuilder = ConditionalRenderingGroupBuilder;
//# sourceMappingURL=conditionalRenderingGroupBuilder.gen.js.map