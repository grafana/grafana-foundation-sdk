"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.RowsLayoutRowBuilder = void 0;
const tslib_1 = require("tslib");
const dashboardv2beta1 = tslib_1.__importStar(require("../dashboardv2beta1"));
class RowsLayoutRowBuilder {
    constructor(title) {
        this.internal = dashboardv2beta1.defaultRowsLayoutRowKind();
        this.internal.kind = "RowsLayoutRow";
        this.internal.spec.title = title;
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    title(title) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultRowsLayoutRowSpec();
        }
        this.internal.spec.title = title;
        return this;
    }
    collapse(collapse) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultRowsLayoutRowSpec();
        }
        this.internal.spec.collapse = collapse;
        return this;
    }
    hideHeader(hideHeader) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultRowsLayoutRowSpec();
        }
        this.internal.spec.hideHeader = hideHeader;
        return this;
    }
    fillScreen(fillScreen) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultRowsLayoutRowSpec();
        }
        this.internal.spec.fillScreen = fillScreen;
        return this;
    }
    conditionalRendering(conditionalRendering) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultRowsLayoutRowSpec();
        }
        const conditionalRenderingResource = conditionalRendering.build();
        this.internal.spec.conditionalRendering = conditionalRenderingResource;
        return this;
    }
    repeat(repeat) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultRowsLayoutRowSpec();
        }
        const repeatResource = repeat.build();
        this.internal.spec.repeat = repeatResource;
        return this;
    }
    gridLayout(gridLayoutKind) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultRowsLayoutRowSpec();
        }
        const gridLayoutKindResource = gridLayoutKind.build();
        this.internal.spec.layout = gridLayoutKindResource;
        return this;
    }
    autoGridLayout(autoGridLayoutKind) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultRowsLayoutRowSpec();
        }
        const autoGridLayoutKindResource = autoGridLayoutKind.build();
        this.internal.spec.layout = autoGridLayoutKindResource;
        return this;
    }
    tabsLayout(tabsLayoutKind) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultRowsLayoutRowSpec();
        }
        const tabsLayoutKindResource = tabsLayoutKind.build();
        this.internal.spec.layout = tabsLayoutKindResource;
        return this;
    }
    rowsLayout(rowsLayoutKind) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultRowsLayoutRowSpec();
        }
        const rowsLayoutKindResource = rowsLayoutKind.build();
        this.internal.spec.layout = rowsLayoutKindResource;
        return this;
    }
}
exports.RowsLayoutRowBuilder = RowsLayoutRowBuilder;
//# sourceMappingURL=rowsLayoutRowBuilder.gen.js.map