"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.ConditionalRenderingGroupSpecBuilder = void 0;
const tslib_1 = require("tslib");
const dashboardv2beta1 = tslib_1.__importStar(require("../dashboardv2beta1"));
class ConditionalRenderingGroupSpecBuilder {
    constructor() {
        this.internal = dashboardv2beta1.defaultConditionalRenderingGroupSpec();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    visibility(visibility) {
        this.internal.visibility = visibility;
        return this;
    }
    condition(condition) {
        this.internal.condition = condition;
        return this;
    }
    items(items) {
        const itemsResources = items.map(builder1 => builder1.build());
        this.internal.items = itemsResources;
        return this;
    }
}
exports.ConditionalRenderingGroupSpecBuilder = ConditionalRenderingGroupSpecBuilder;
//# sourceMappingURL=conditionalRenderingGroupSpecBuilder.gen.js.map