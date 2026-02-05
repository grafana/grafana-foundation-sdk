"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.TabsLayoutTabSpecBuilder = void 0;
const tslib_1 = require("tslib");
const dashboardv2beta1 = tslib_1.__importStar(require("../dashboardv2beta1"));
class TabsLayoutTabSpecBuilder {
    constructor() {
        this.internal = dashboardv2beta1.defaultTabsLayoutTabSpec();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    title(title) {
        this.internal.title = title;
        return this;
    }
    layout(layout) {
        const layoutResource = layout.build();
        this.internal.layout = layoutResource;
        return this;
    }
    conditionalRendering(conditionalRendering) {
        const conditionalRenderingResource = conditionalRendering.build();
        this.internal.conditionalRendering = conditionalRenderingResource;
        return this;
    }
    repeat(repeat) {
        const repeatResource = repeat.build();
        this.internal.repeat = repeatResource;
        return this;
    }
}
exports.TabsLayoutTabSpecBuilder = TabsLayoutTabSpecBuilder;
//# sourceMappingURL=tabsLayoutTabSpecBuilder.gen.js.map