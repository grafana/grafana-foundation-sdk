"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.RowsLayoutRowSpecBuilder = void 0;
const tslib_1 = require("tslib");
const dashboardv2beta1 = tslib_1.__importStar(require("../dashboardv2beta1"));
class RowsLayoutRowSpecBuilder {
    constructor() {
        this.internal = dashboardv2beta1.defaultRowsLayoutRowSpec();
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
    collapse(collapse) {
        this.internal.collapse = collapse;
        return this;
    }
    hideHeader(hideHeader) {
        this.internal.hideHeader = hideHeader;
        return this;
    }
    fillScreen(fillScreen) {
        this.internal.fillScreen = fillScreen;
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
    layout(layout) {
        const layoutResource = layout.build();
        this.internal.layout = layoutResource;
        return this;
    }
}
exports.RowsLayoutRowSpecBuilder = RowsLayoutRowSpecBuilder;
//# sourceMappingURL=rowsLayoutRowSpecBuilder.gen.js.map