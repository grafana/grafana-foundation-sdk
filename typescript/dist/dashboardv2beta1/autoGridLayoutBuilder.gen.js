"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.AutoGridLayoutBuilder = void 0;
const tslib_1 = require("tslib");
const dashboardv2beta1 = tslib_1.__importStar(require("../dashboardv2beta1"));
class AutoGridLayoutBuilder {
    constructor() {
        this.internal = dashboardv2beta1.defaultAutoGridLayoutKind();
        this.internal.kind = "AutoGridLayout";
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    maxColumnCount(maxColumnCount) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultAutoGridLayoutSpec();
        }
        this.internal.spec.maxColumnCount = maxColumnCount;
        return this;
    }
    columnWidthMode(columnWidthMode) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultAutoGridLayoutSpec();
        }
        this.internal.spec.columnWidthMode = columnWidthMode;
        return this;
    }
    columnWidth(columnWidth) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultAutoGridLayoutSpec();
        }
        this.internal.spec.columnWidth = columnWidth;
        return this;
    }
    rowHeightMode(rowHeightMode) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultAutoGridLayoutSpec();
        }
        this.internal.spec.rowHeightMode = rowHeightMode;
        return this;
    }
    rowHeight(rowHeight) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultAutoGridLayoutSpec();
        }
        this.internal.spec.rowHeight = rowHeight;
        return this;
    }
    fillScreen(fillScreen) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultAutoGridLayoutSpec();
        }
        this.internal.spec.fillScreen = fillScreen;
        return this;
    }
    items(items) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultAutoGridLayoutSpec();
        }
        const itemsResources = items.map(builder1 => builder1.build());
        this.internal.spec.items = itemsResources;
        return this;
    }
    item(item) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultAutoGridLayoutSpec();
        }
        if (!this.internal.spec.items) {
            this.internal.spec.items = [];
        }
        const itemResource = item.build();
        this.internal.spec.items.push(itemResource);
        return this;
    }
}
exports.AutoGridLayoutBuilder = AutoGridLayoutBuilder;
//# sourceMappingURL=autoGridLayoutBuilder.gen.js.map