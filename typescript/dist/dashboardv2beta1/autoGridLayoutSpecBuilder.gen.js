"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.AutoGridLayoutSpecBuilder = void 0;
const tslib_1 = require("tslib");
const dashboardv2beta1 = tslib_1.__importStar(require("../dashboardv2beta1"));
class AutoGridLayoutSpecBuilder {
    constructor() {
        this.internal = dashboardv2beta1.defaultAutoGridLayoutSpec();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    maxColumnCount(maxColumnCount) {
        this.internal.maxColumnCount = maxColumnCount;
        return this;
    }
    columnWidthMode(columnWidthMode) {
        this.internal.columnWidthMode = columnWidthMode;
        return this;
    }
    columnWidth(columnWidth) {
        this.internal.columnWidth = columnWidth;
        return this;
    }
    rowHeightMode(rowHeightMode) {
        this.internal.rowHeightMode = rowHeightMode;
        return this;
    }
    rowHeight(rowHeight) {
        this.internal.rowHeight = rowHeight;
        return this;
    }
    fillScreen(fillScreen) {
        this.internal.fillScreen = fillScreen;
        return this;
    }
    items(items) {
        const itemsResources = items.map(builder1 => builder1.build());
        this.internal.items = itemsResources;
        return this;
    }
}
exports.AutoGridLayoutSpecBuilder = AutoGridLayoutSpecBuilder;
//# sourceMappingURL=autoGridLayoutSpecBuilder.gen.js.map